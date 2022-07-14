package providers

import (
	"context"
	"fmt"
	"log"
	"net"

	censorPb "censor.services/app/domain/proto"

	"google.golang.org/grpc"

	"github.com/spf13/viper"

	"go.uber.org/fx"
)

func StartGrpcService(lifecycle fx.Lifecycle, srv censorPb.CensorGrpcServiceServer, config *viper.Viper) {
	listenGrpc := config.GetString("grpc.ListenHTTP")

	lis, err := net.Listen("tcp", listenGrpc)
	if nil != err {
		return
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	censorPb.RegisterCensorGrpcServiceServer(grpcServer, srv)

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("lifecycle.OnStart StartGrpcService on :", listenGrpc)
			go func() {
				err = grpcServer.Serve(lis)
				if err != nil {
					fmt.Println("grpcServer.Serve err : ", err)
				}
			}()
			return nil
		},
	})
}
