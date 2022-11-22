package repository

import (
	"context"
	"github.com/cvetkovski98/zvax-users/internal/users"
	"github.com/cvetkovski98/zvax-users/internal/users/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgUserRepository struct {
	pool *pgxpool.Pool
}

func (repository *PgUserRepository) FindByID(ctx context.Context, userId int64) (*model.User, error) {
	var user = new(model.User)
	err := repository.pool.QueryRow(
		ctx, "select * from users where id = $1", userId,
	).Scan(&user.UserId, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindAll returns all users from the database.
func (repository *PgUserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	var userList []model.User
	rows, err := repository.pool.Query(ctx, "select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user = new(model.User)
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		userList = append(userList, *user)
	}
	return userList, nil
}

// InsertOne creates a new user in the database.
func (repository *PgUserRepository) InsertOne(ctx context.Context, user *model.User) (*model.User, error) {
	err := repository.pool.QueryRow(
		ctx, "insert into users (name, email, phone) values ($1, $2, $3) returning id, name, email, phone",
		user.Name, user.Email, user.Phone,
	).Scan(&user.UserId, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewPgUserRepository(pool *pgxpool.Pool) users.UserRepository {
	return &PgUserRepository{pool: pool}
}
