package auth

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/drewspitsin/auth/internal/client/db"
	"github.com/drewspitsin/auth/internal/model"
	"github.com/drewspitsin/auth/internal/repository"
	"github.com/drewspitsin/auth/internal/repository/auth/converter"
	modelRepo "github.com/drewspitsin/auth/internal/repository/auth/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	id          = "id"
	table       = "user_table"
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

func (s *repo) Create(ctx context.Context, info *model.User) (int64, error) {

	builderInsert := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(username, email, password).
		Values(info.Name, info.Email, info.Password).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "auth_repository.Create",
		QueryRaw: query,
	}

	var userTableID int64
	err = s.db.DB().QueryRowContext(ctx, q, args...).Scan(&userTableID)
	if err != nil {
		return 0, err
	}

	return userTableID, nil
}

func (s *repo) Get(ctx context.Context, userTableID int64) (*model.User, error) {
	// Делаем запрос на получение измененной записи из таблицы user_table
	builderSelectOne := sq.Select(id, username, email, role, createdAtPg, updatedAtPg).
		From(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{id: userTableID}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
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

func (s *repo) Update(ctx context.Context, info *model.User) error {

	var updatedAt sql.NullTime
	var updatedAtTime *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtTime = timestamppb.New(updatedAt.Time)
	}

	builderUpdate := sq.Update(table).
		PlaceholderFormat(sq.Dollar).
		Set(username, info.Name).
		Set(email, info.Email).
		Set(role, info.Role).
		Set(updatedAtPg, updatedAtTime).
		Where(sq.Eq{id: info.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v ", err)
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Update",
		QueryRaw: query,
	}

	res, err := s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Fatalf("failed to update user: %v tag: %v", err, res)
		return err
	}

	return nil
}

func (s *repo) Delete(ctx context.Context, info *model.User) error {
	builderInsert := sq.Delete(table).
		Where(sq.Eq{id: info.ID}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Delete",
		QueryRaw: query,
	}
	res, err := s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Fatalf("failed to delete user: %v tag: %v", err, res)
		return err
	}

	return nil
}
