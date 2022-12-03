package repository

import (
	"context"
	"errors"
	"sync"

	keys "github.com/cvetkovski98/zvax-keys/internal"
	"github.com/cvetkovski98/zvax-keys/internal/model"
)

type mock struct {
	db   map[int]*model.Key
	lock sync.RWMutex
}

func (m *mock) FindAllByHolder(ctx context.Context, holder string) ([]*model.Key, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var keys []*model.Key
	for _, k := range m.db {
		if k.Holder == holder {
			keys = append(keys, k)
		}
	}
	return keys, nil
}

func (m *mock) FindOneById(ctx context.Context, id int) (*model.Key, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if k, ok := m.db[id]; ok {
		return k, nil
	}
	return nil, errors.New("not found")
}

func (m *mock) InsertOne(ctx context.Context, key *model.Key) (*model.Key, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	key.Id = len(m.db) + 1
	m.db[key.Id] = key
	return key, nil
}

func NewMockKeyRepository() keys.Repository {
	return &mock{
		db:   make(map[int]*model.Key),
		lock: sync.RWMutex{},
	}
}
