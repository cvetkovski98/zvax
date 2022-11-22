package repository

import (
	"context"
	"fmt"
	"github.com/cvetkovski98/zvax-slots/internal"
	"github.com/cvetkovski98/zvax-slots/internal/model"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"time"
)

const slotsOrderedSetKey = "slots"
const unconfirmedReservationSetKey = "unconfirmed_reservations"
const confirmationTimeout = time.Minute * 5

type RedisSlotRepository struct {
	rdb *redis.Client
}

// FindOneByKey returns a slot by key
func (repository *RedisSlotRepository) FindOneByKey(
	ctx context.Context,
	key string,
) (*model.Slot, error) {
	result, err := repository.rdb.HGetAll(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error getting slot with key=%s",
			key,
		))
	}
	return model.NewSlotFromMap(result)
}

func (repository *RedisSlotRepository) InsertOne(
	ctx context.Context,
	slot *model.Slot,
) (*model.Slot, error) {
	key := model.NewSlotRedisId(slot.Location, slot.DateTime)
	slot.SlotID = &key
	_, err := repository.rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		err := pipe.HSet(ctx, key, slot.ToMap()).Err()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf(
				"error setting key=%s to value=%v",
				key,
				slot,
			))
		}
		err = pipe.ZAdd(ctx, slotsOrderedSetKey, redis.Z{
			Score:  float64(slot.DateTime.Unix()),
			Member: key,
		}).Err()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf(
				"error adding slot with key=%s to set=%s",
				key,
				slotsOrderedSetKey,
			))
		}
		return nil
	})
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error inserting slot with key=%s",
			key,
		))
	}
	return slot, nil
}

func (repository *RedisSlotRepository) FindAllWithDateTimeBetween(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]*model.Slot, error) {
	slotIds, err := repository.rdb.ZRangeByScore(ctx, slotsOrderedSetKey, &redis.ZRangeBy{
		Min:    strconv.FormatInt(from.Unix(), 10),
		Max:    strconv.FormatInt(to.Unix(), 10),
		Offset: 0,
		Count:  -1,
	}).Result()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error getting slots in time range from=%s to=%s",
			from,
			to,
		))
	}
	commands, err := repository.rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, slotId := range slotIds {
			err := pipe.HGetAll(ctx, slotId).Err()
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf(
					"error getting slot with id=%s",
					slotId,
				))
			}
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error getting slots in time range from=%s to=%s",
			from,
			to,
		))
	}
	slots := make([]*model.Slot, 0, len(commands))
	for i, command := range commands {
		sMap, err := command.(*redis.MapStringStringCmd).Result()
		if err != nil && err != redis.Nil {
			return nil, errors.Wrap(err, fmt.Sprintf(
				"error processing command result for slot=%s",
				slotIds[i],
			))
		}
		slot, err := model.NewSlotFromMap(sMap)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf(
				"error creating slot from map for slot=%s",
				slotIds[i],
			))
		}
		if slot.Available {
			slots = append(slots, slot)
		}
	}
	return slots, nil
}

func (repository *RedisSlotRepository) ReserveOneByKey(
	ctx context.Context,
	key string,
) (*model.Reservation, error) {
	reservationId := model.NewReservationRedisId(key)
	tfx := func(tx *redis.Tx) error {
		slot, err := repository.FindOneByKey(ctx, key)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf(
				"error finding slot with key=%s",
				key,
			))
		}
		if !slot.Available {
			return errors.New("slot is not available")
		}
		// update the slot to unavailable
		err = tx.HSet(ctx, *slot.SlotID, "available", "false").Err()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf(
				"error setting key=%s to value=%v",
				*slot.SlotID,
				slot,
			))
		}
		// create a reservation for the slot
		validity := time.Now().Add(confirmationTimeout)
		reservation := &model.Reservation{
			ReservationID: &reservationId,
			Slot:          slot,
			ValidUntil:    validity,
		}
		err = tx.HSet(ctx, reservationId, reservation.ToMap()).Err()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf(
				"error setting key=%s to value=%v",
				reservationId,
				reservation,
			))
		}
		err = tx.ZAdd(ctx, unconfirmedReservationSetKey, redis.Z{
			Score:  float64(validity.Unix()),
			Member: reservationId,
		}).Err()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf(
				"error adding key=%s to set=%s",
				reservationId,
				unconfirmedReservationSetKey,
			))
		}
		return nil
	}
	err := repository.rdb.Watch(ctx, tfx, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error reserving slot with key=%s",
			key,
		))
	}
	return repository.findReservationById(ctx, reservationId)
}

func (repository *RedisSlotRepository) findReservationById(
	ctx context.Context,
	reservationId string,
) (*model.Reservation, error) {
	rMap, err := repository.rdb.HGetAll(ctx, reservationId).Result()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error getting reservation with key=%s",
			reservationId,
		))
	}
	sMap, err := repository.rdb.HGetAll(ctx, rMap["slotId"]).Result()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error getting slot with key=%s",
			rMap["slotId"],
		))
	}
	reservation, err := model.NewReservationFromMaps(rMap, sMap)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error creating reservation from map for key=%s",
			reservationId,
		))
	}
	return reservation, nil
}

func (repository *RedisSlotRepository) ConfirmOneByReservationId(ctx context.Context, reservationId string) (*model.Reservation, error) {
	reservation, err := repository.findReservationById(ctx, reservationId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error finding reservation with id=%s",
			reservationId,
		))
	}
	tfx := func(tx *redis.Tx) error {
		// check if the reservation is in unconfirmed set
		exists, err := tx.SIsMember(ctx, unconfirmedReservationSetKey, reservationId).Result()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf(
				"error checking if key=%s is in set=%s",
				reservationId,
				unconfirmedReservationSetKey,
			))
		}
		if exists {
			// remove it
			err = tx.SRem(ctx, unconfirmedReservationSetKey, reservationId).Err()
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf(
					"error removing key=%s from set=%s",
					reservationId,
					unconfirmedReservationSetKey,
				))
			}
		} else {
			// check if the slot is available
			if !reservation.Slot.Available {
				return errors.New(fmt.Sprintf(
					"slot with key=%s is not available",
					*reservation.Slot.SlotID,
				))
			}
			// try setting the slot to unavailable
			err = tx.HSet(ctx, *reservation.Slot.SlotID, "available", "false").Err()
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf(
					"error setting key=%s to value=%v",
					*reservation.Slot.SlotID,
					reservation.Slot,
				))
			}
		}
		return nil
	}
	err = repository.rdb.Watch(ctx, tfx, reservationId, *reservation.Slot.SlotID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"error confirming reservation with id=%s",
			reservationId,
		))
	}
	return reservation, nil
}

// CleanUpExpiredReservations finds and deletes expired reservations and sets their slots to available
func (repository *RedisSlotRepository) CleanUpExpiredReservations(ctx context.Context) error {
	zRangeToNow := &redis.ZRangeBy{
		Min:    "-inf",
		Max:    fmt.Sprintf("%d", time.Now().Unix()),
		Offset: 0,
		Count:  -1,
	}
	expiredReservationIds, err := repository.rdb.ZRangeByScore(
		ctx,
		unconfirmedReservationSetKey,
		zRangeToNow,
	).Result()
	if err != nil {
		return errors.Wrap(err, "error getting expired reservations")
	}

	tfx := func(tx *redis.Tx) error {
		fetchReservationsPipe := tx.TxPipeline()
		for _, reservationId := range expiredReservationIds {
			fetchReservationsPipe.HGetAll(ctx, reservationId)
		}
		rCmds, err := fetchReservationsPipe.Exec(ctx)
		if err != nil {
			return errors.Wrap(err, "error executing fetch reservations pipe")
		}
		setSlotsAvailablePipe := tx.TxPipeline()
		for _, rCmd := range rCmds {
			rMap, err := rCmd.(*redis.MapStringStringCmd).Result()
			if err != nil {
				return errors.Wrap(err, "error getting reservation map")
			}
			if sId, ok := rMap["slotId"]; ok {
				setSlotsAvailablePipe.HSet(ctx, sId, "available", "true")
			}
		}
		if _, err = setSlotsAvailablePipe.Exec(ctx); err != nil {
			return errors.Wrap(err, "error executing set slots available pipe")
		}
		if _, err := tx.ZRemRangeByScore(
			ctx,
			unconfirmedReservationSetKey,
			zRangeToNow.Min,
			zRangeToNow.Max,
		).Result(); err != nil {
			return errors.Wrap(err, "error removing expired reservations")
		}
		return nil
	}
	if err := repository.rdb.Watch(ctx, tfx, expiredReservationIds...); err != nil {
		return errors.Wrap(err, "error watching expired reservations")
	} else {
		log.Printf(
			"removed %d expired reservations from set=%s",
			len(expiredReservationIds),
			unconfirmedReservationSetKey,
		)
		return nil
	}
}

func NewRedisSlotRepository(rdb *redis.Client) internal.SlotRepository {
	return &RedisSlotRepository{rdb: rdb}
}
