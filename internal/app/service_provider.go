package app

import (
	"context"
	"log"

	"github.com/drewspitsin/auth/internal/api/access"
	"github.com/drewspitsin/auth/internal/api/auth"
	"github.com/drewspitsin/auth/internal/api/login"
	"github.com/drewspitsin/auth/internal/client/db"
	"github.com/drewspitsin/auth/internal/client/db/pg"
	"github.com/drewspitsin/auth/internal/client/db/transaction"
	"github.com/drewspitsin/auth/internal/closer"
	"github.com/drewspitsin/auth/internal/config"
	"github.com/drewspitsin/auth/internal/config/env"
	"github.com/drewspitsin/auth/internal/repository"
	accessRepository "github.com/drewspitsin/auth/internal/repository/access"
	authRepository "github.com/drewspitsin/auth/internal/repository/auth"
	loginRepository "github.com/drewspitsin/auth/internal/repository/login"
	"github.com/drewspitsin/auth/internal/service"
	accessService "github.com/drewspitsin/auth/internal/service/access"
	authService "github.com/drewspitsin/auth/internal/service/auth"
	loginService "github.com/drewspitsin/auth/internal/service/login"
)

// func main() {
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	reflection.Register(s)
// 	descAuth.RegisterAuthV1Server(s, &serverAuth{})
// 	descAccess.RegisterAccessV1Server(s, &serverAccess{})

// 	log.Printf("server listening at %v", lis.Addr())

// 	if err = s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	dbClient         db.Client
	txManager        db.TxManager
	authRepository   repository.AuthRepository
	loginRepository  repository.LoginRepository
	accessRepository repository.AccessRepository

	authService   service.AuthService
	loginService  service.LoginService
	accessService service.AccessService

	authImpl   *auth.Implementation
	loginImpl  *login.Implementation
	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.DBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.AuthRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) LoginRepository(ctx context.Context) repository.LoginRepository {
	if s.loginRepository == nil {
		s.loginRepository = loginRepository.NewRepository(s.DBClient(ctx))
	}

	return s.loginRepository
}

func (s *serviceProvider) LoginService(ctx context.Context) service.LoginService {
	if s.loginService == nil {
		s.loginService = loginService.NewService(
			s.LoginRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.loginService
}

func (s *serviceProvider) LoginImpl(ctx context.Context) *login.Implementation {
	if s.LoginImpl == nil {
		s.loginImpl = login.NewImplementation(s.LoginService(ctx))
	}

	return s.loginImpl
}

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.DBClient(ctx))
	}

	return s.accessRepository
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.AccessRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.accessService
}
func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}
