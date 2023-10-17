package grpc_server

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *UserV1Server) CreatePg(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	builderInsert := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(username, email, password).
		Values(req.Name, req.Email, req.Password).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var userTableID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&userTableID)
	if err != nil {
		log.Fatalf("failed to insert user_table: %v", err)
	}

	log.Printf("inserted user_table with id: %d", userTableID)

	return &desc.CreateResponse{
		Id: userTableID,
	}, nil
}

func (s *UserV1Server) GetPg(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	// Делаем запрос на получение измененной записи из таблицы user_table
	builderSelectOne := sq.Select(id, username, email, role, createdAtPg, updatedAtPg).
		From(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{id: req.GetId()}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var id int64
	var user_name, email string
	var role int32
	var createdAt time.Time
	var updatedAt sql.NullTime

	err = s.pool.QueryRow(ctx, query, args...).Scan(&id, &user_name, &email, &role, &createdAt, &updatedAt)
	if err != nil {
		log.Fatalf("failed to select user_tables: %v", err)
	}

	log.Printf("id: %d, user_name: %s, email: %s, role: %v, created_at: %v, updated_at: %v\n", id, user_name, email, role, createdAt, updatedAt)

	var updatedAtTime *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtTime = timestamppb.New(updatedAt.Time)
	}

	return &desc.GetResponse{
		Id:        id,
		Name:      user_name,
		Email:     email,
		Role:      desc.Role(role),
		CreatedAt: &timestamppb.Timestamp{},
		UpdatedAt: updatedAtTime,
	}, nil
}

func (s *UserV1Server) UpdatePg(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {

	var updatedAt sql.NullTime
	var updatedAtTime *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtTime = timestamppb.New(updatedAt.Time)
	}

	builderUpdate := sq.Update(table).
		PlaceholderFormat(sq.Dollar).
		Set(username, req.Name).
		Set(email, req.Email).
		Set(role, req.Role).
		Set(updatedAtPg, updatedAtTime).
		Where(sq.Eq{id: req.Id})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
		return nil, err
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to update note: %v", err)
		return nil, err
	}

	log.Printf("updated %d rows", res.RowsAffected())
	return &emptypb.Empty{}, nil
}

func (s *UserV1Server) DeletePg(ctx context.Context, req *desc.GetRequest) (*empty.Empty, error) {
	builderInsert := sq.Delete(table).
		Where(sq.Eq{id: req.GetId()}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
		return nil, err
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to update note: %v tag: %v", err, res)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
