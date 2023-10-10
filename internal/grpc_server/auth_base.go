package grpc_server

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// id serial primary key,
// UserName text not null,
// Email text not null,
// Pswd text not null,
// created_at timestamp not null default now(),
// updated_at timestamp

func (s *UserV1Server) CreatePg(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	builderInsert := sq.Insert("user_table").
		PlaceholderFormat(sq.Dollar).
		Columns("user_name", "email", "pswd").
		Values(gofakeit.City(), gofakeit.Address().Zip, gofakeit.City()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var user_tableID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&user_tableID)
	if err != nil {
		log.Fatalf("failed to insert user_table: %v", err)
	}

	log.Printf("inserted user_table with id: %d", user_tableID)

	return &desc.CreateResponse{
		Id: user_tableID,
	}, nil
}

func (s *UserV1Server) GetPg(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	// Делаем запрос на получение измененной записи из таблицы user_table
	builderSelectOne := sq.Select("id", "user_name", "email", "role", "created_at", "updated_at").
		From("user_table").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
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

func (s *UserV1Server) UpdatePg(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	return nil, nil
}

func (s *UserV1Server) DeletePg(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	return nil, nil
}
