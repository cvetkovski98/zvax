package repository

import (
	"context"
	"github.com/cvetkovski98/zvax-keys/internal"
	"github.com/cvetkovski98/zvax-keys/internal/model"
	"github.com/pkg/errors"
	"sync"
)

type memKeyRepository struct {
	keys map[int64]*model.Key
	lock sync.RWMutex
}

func (repository *memKeyRepository) InsertOne(ctx context.Context, key *model.Key) (*model.Key, error) {
	repository.lock.Lock()
	defer repository.lock.Unlock()
	id := int64(len(repository.keys))
	key.KeyId = &id
	repository.keys[id] = key
	return key, nil
}

func (repository *memKeyRepository) FindAll(ctx context.Context) ([]*model.Key, error) {
	repository.lock.RLock()
	defer repository.lock.RUnlock()
	keys := make([]*model.Key, 0)
	for _, key := range repository.keys {
		keys = append(keys, key)
	}
	return keys, nil
}

func (repository *memKeyRepository) FindOneById(ctx context.Context, keyId int64) (*model.Key, error) {
	repository.lock.RLock()
	defer repository.lock.RUnlock()
	key, ok := repository.keys[keyId]
	if !ok {
		return nil, errors.New("key not found")
	}
	return key, nil
}

func NewInMemoryKeyRepository(source map[int64]*model.Key) internal.KeyRepository {
	return &memKeyRepository{keys: source, lock: sync.RWMutex{}}
}
