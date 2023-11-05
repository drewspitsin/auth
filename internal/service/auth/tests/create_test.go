package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"

	"github.com/drewspitsin/auth/internal/client/db"
	txMocks "github.com/drewspitsin/auth/internal/client/db/mocks"
	"github.com/drewspitsin/auth/internal/client/db/transaction"
	"github.com/drewspitsin/auth/internal/model"
	"github.com/drewspitsin/auth/internal/repository"
	repoMocks "github.com/drewspitsin/auth/internal/repository/mocks"
	"github.com/drewspitsin/auth/internal/service/auth"
)

type TxMock struct {
	pgxpool.Tx
}

func (t *TxMock) Commit(_ context.Context) error {
	return nil
}

func (t *TxMock) Rollback(_ context.Context) error {
	return nil
}

func TestCreate(t *testing.T) {
	t.Parallel()
	type authRepositoryMockFunc func(mc *minimock.Controller) repository.AuthRepository
	type txTransactorMockFunc func(mc *minimock.Controller) db.Transactor
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.UserCreate
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Animal()
		email     = gofakeit.Email()
		password  = gofakeit.Animal()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		repoErr = fmt.Errorf("can't begin transaction: repo error")

		req = &model.UserCreate{
			UserUpdate: model.UserUpdate{ID: id, Name: name, Email: email, Role: 0},
			Password:   password,
		}

		res = id

		txM TxMock

		resGet = &model.User{
			UserCreate: model.UserCreate{UserUpdate: model.UserUpdate{ID: id, Name: name, Email: email, Role: 0}, Password: password},
			CreatedAt:  createdAt,
			UpdatedAt:  sql.NullTime{Time: updatedAt, Valid: true},
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		authRepositoryMock authRepositoryMockFunc
		txTransactorMock   txTransactorMockFunc
		txManagerMock      txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				mock.CreateMock.Return(res, nil)
				mock.GetMock.Return(resGet, nil)
				return mock
			},
			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := txMocks.NewTransactorMock(mc)
				mock.BeginTxMock.Return(&txM, nil)
				return mock
			},
		},

		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			authRepositoryMock: func(mc *minimock.Controller) repository.AuthRepository {
				mock := repoMocks.NewAuthRepositoryMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.name == "service error case" {
				authRepoMock := tt.authRepositoryMock(mc)
				tx := tt.txManagerMock(mc)
				service := auth.NewService(authRepoMock, tx)
				newID, err := service.Create(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, newID)
			} else {
				authRepoMock := tt.authRepositoryMock(mc)
				txTransact := transaction.NewTransactionManager(tt.txTransactorMock(mc))
				service := auth.NewService(authRepoMock, txTransact)
				newID, err := service.Create(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, newID)
			}

		})
	}
}
