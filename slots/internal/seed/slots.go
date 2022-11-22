package main

import (
	"context"
	"github.com/cvetkovski98/zvax-slots/internal/model"
	"github.com/cvetkovski98/zvax-slots/internal/repository"
	"github.com/cvetkovski98/zvax-slots/internal/service"
	"github.com/cvetkovski98/zvax-slots/pkg/redis"
	"log"
	"time"
)

func GenerateSlots(startDate time.Time, endDate time.Time) ([]model.Slot, error) {
	var slots = make([]model.Slot, 0)
	currentDate := time.Date(
		startDate.Year(),
		startDate.Month(),
		startDate.Day(),
		0, 0, 0, 0,
		startDate.Location(),
	)
	for currentDate.Before(endDate) {
		if currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday {
			currentDate = currentDate.Add(time.Hour * 24)
			continue
		}
		startOfWorkDay := time.Date(
			currentDate.Year(),
			currentDate.Month(),
			currentDate.Day(),
			9, 0, 0, 0,
			currentDate.Location(),
		)
		endOfWorkDay := time.Date(
			currentDate.Year(),
			currentDate.Month(),
			currentDate.Day(),
			18, 0, 0, 0,
			currentDate.Location(),
		)
		for startOfWorkDay.Before(endOfWorkDay) {
			//available := rand.Intn(100) < 40
			slots = append(slots, model.Slot{
				SlotID:    nil,
				DateTime:  startOfWorkDay,
				Location:  "random",
				Available: true,
			})
			startOfWorkDay = startOfWorkDay.Add(time.Minute * 30)
		}
		currentDate = currentDate.Add(time.Hour * 24)
	}
	return slots, nil
}

func main() {
	var (
		start = time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
		end   = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	)
	rdb, err := redis.NewRedisConn()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	log.Printf("connected to db")
	slotRepository := repository.NewRedisSlotRepository(rdb)
	slotService := service.NewSlotServiceImpl(slotRepository)
	generated, err := GenerateSlots(start, end)
	if err != nil {
		log.Fatalf("failed to generate slots: %v", err)
	}
	log.Printf("generated %d slots", len(generated))
	inserted := 0
	for _, slot := range generated {
		_, err := slotService.CreateSlot(context.Background(), &slot)
		if err != nil {
			log.Fatalf("failed to create slot: %v", err)
		}
		inserted++
	}
	log.Printf("created %d slots", inserted)
	slots, err := slotService.GetSlotsAtLocationBetween(context.Background(), &model.PageRequest{
		StartDate: start,
		EndDate:   end,
		Location:  "random",
	})
	if err != nil {
		log.Fatalf("failed to get slots: %v", err)
	}
	log.Printf("got %d slots", len(slots.Items))
	_, err = slotService.CreateReservation(context.Background(), "slot:random:2021-01-01:09-00")
	_, err = slotService.CreateReservation(context.Background(), "slot:random:2021-01-01:09-30")
	_, err = slotService.CreateReservation(context.Background(), "slot:random:2021-01-01:10-00")
	_, err = slotService.CreateReservation(context.Background(), "slot:random:2021-01-01:10-30")
	if err != nil {
		log.Fatalf("failed to get slot: %v", err)
	}
	go func() {
		for {
			err := slotRepository.CleanUpExpiredReservations(context.Background())
			if err != nil {
				log.Fatalf("failed to clean up expired reservations: %v", err)
			}
			time.Sleep(time.Second * 10)
		}
	}()
	select {}
}
