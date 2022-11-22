package test

import (
	"context"
	"github.com/cvetkovski98/zvax-users/internal/users"
	"github.com/cvetkovski98/zvax-users/internal/users/model"
	"github.com/pkg/errors"
	"sync"
)

type MockUserRepository struct {
	users map[int64]*model.User
	lock  sync.RWMutex
}

func (repository *MockUserRepository) FindByID(ctx context.Context, userId int64) (*model.User, error) {
	repository.lock.RLock()
	defer repository.lock.RUnlock()
	user, ok := repository.users[userId]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (repository *MockUserRepository) InsertOne(ctx context.Context, user *model.User) (*model.User, error) {
	repository.lock.Lock()
	defer repository.lock.Unlock()
	id := int64(len(repository.users) + 1)
	user.UserId = &id
	repository.users[*user.UserId] = user
	return user, nil
}

func (repository *MockUserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	repository.lock.RLock()
	defer repository.lock.RUnlock()
	users := make([]model.User, 0)
	for _, user := range repository.users {
		users = append(users, *user)
	}
	return users, nil
}

func NewMockUserRepository(source map[int64]*model.User) users.UserRepository {
	return &MockUserRepository{
		users: source,
		lock:  sync.RWMutex{},
	}
}
