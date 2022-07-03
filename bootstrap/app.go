package bootstrap

import (
	"censor.services/app/domain/repositories"
	"censor.services/app/domain/services"
	"censor.services/app/grpc"
	"censor.services/app/http/controllers"
	"censor.services/app/http/middlewares"
	"censor.services/app/providers"
	"censor.services/configs"
	"censor.services/routes"
	"github.com/henrion-y/base.services/database/mongo"
	middlewares2 "github.com/henrion-y/base.services/http/gin/middlewares"
	"github.com/henrion-y/base.services/infra/censor/baiduCensor"
	"github.com/henrion-y/base.services/infra/jwt"
	"go.uber.org/fx"
)

func App() *fx.App {
	options := make([]fx.Option, 0)

	comOptions := []fx.Option{
		// Configurations (./config)
		fx.Provide(configs.NewConfigProvider),

		// Providers (./app/providers)
		fx.Provide(providers.NewGinProvider),
		fx.Provide(providers.NewLoggerProvider),

		fx.Provide(baiduCensor.NewBaiduCensorClient),
		fx.Provide(jwt.NewJWTService),

		// Providers (db)
		fx.Provide(mongo.NewDbProvider),

		// Middlewares (./app/http/middlewares)
		fx.Provide(middlewares2.NewJWTAuthMiddleware),
		fx.Provide(middlewares.NewLogMiddleware),

		// NerRepository
		// Repositories (./app/domain/repositories)
		fx.Provide(repositories.NewSensitiveWordRepository),

		// NewService
		// Services (./app/domain/services)
		fx.Provide(services.NewCensorService),

		// NewController
		// Controllers (./app/http/controllers)
		fx.Provide(controllers.NewCensorController),
		fx.Provide(grpc.NewCensorGrpcService),

		fx.Invoke(middlewares.GlobalMiddlewares),
		fx.Invoke(providers.StartGrpcService),
		fx.Invoke(routes.APIRoutes),
		fx.Invoke(providers.StartService),
	}

	options = append(options, comOptions...)

	return fx.New(options...)
}
