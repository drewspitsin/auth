package auth

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/drewspitsin/auth/internal/client/db"
	"github.com/drewspitsin/auth/internal/model"
	"github.com/drewspitsin/auth/internal/repository"
	"github.com/drewspitsin/auth/internal/repository/auth/converter"
	modelRepo "github.com/drewspitsin/auth/internal/repository/auth/model"
)

const (
	id          = "id"
	table       = "users"
	username    = "username"
	email       = "email"
	password    = "password"
	role        = "role"
	createdAtPg = "created_at"
	updatedAtPg = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(dbClient db.Client) repository.AuthRepository {
	return &repo{db: dbClient}
}

func (s *repo) Create(ctx context.Context, info *model.UserCreate) (int64, error) {
	builderInsert := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(username, email, password, role).
		Values(info.Name, info.Email, info.Password, info.Role).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "auth_repository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = s.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// Делаем запрос на получение измененной записи из таблицы user_table
func (s *repo) Get(ctx context.Context, userID int64) (*model.User, error) {
	builderSelectOne := sq.Select(id, username, email, role, createdAtPg, updatedAtPg).
		From(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{id: userID}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "auth_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = s.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (s *repo) Update(ctx context.Context, info *model.UserUpdate) error {
	builderUpdate := sq.Update(table).
		PlaceholderFormat(sq.Dollar).
		Set(username, info.Name).
		Set(email, info.Email).
		Set(role, info.Role).
		Set(updatedAtPg, time.Now()).
		Where(sq.Eq{id: info.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %v ", err)
	}

	q := db.Query{
		Name:     "auth_repository.Update",
		QueryRaw: query,
	}

	res, err := s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %v tag: %v", err, res)
	}

	return nil
}

func (s *repo) Delete(ctx context.Context, id_d int64) error {
	builderInsert := sq.Delete(table).
		Where(sq.Eq{id: id_d}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builderInsert.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "auth_repository.Delete",
		QueryRaw: query,
	}
	res, err := s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v tag: %v", err, res)
	}

	return nil
}
