package app

import (
	"net"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/auth/proto/auth"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/account/profile/domain/profile"
	"google.golang.org/grpc"
)

type GrpcServer interface {
	auth.AuthServer
	Run() errors.Error
	Stop()
}

type grpcServer struct {
	auth.UnsafeAuthServer
	account.GrpcAccountsGetter
	account.GrpcAccountGetter
	profile.GrpcProfilesGetter
	profile.GrpcProfileGetter
	logger log.Logger
	config struct {
		ListenAddr string
	}
	// server
	gRpcServer *grpc.Server
}

func (a *App[acc, prof, regReq, reg]) newGrpc(registry registry.Registry) GrpcServer {
	s := &grpcServer{
		logger:             a.Logger.WithTrace("grpc"),
		GrpcAccountsGetter: a.Account.GrpcListHandler,
		GrpcAccountGetter:  a.Account.GrpcGetHandler,
		GrpcProfilesGetter: a.Account.Profile.GrpcListHandler,
		GrpcProfileGetter:  a.Account.Profile.GrpcGetHandler,
	}
	if err := registry.Unmarshal(&s.config); err != nil {
		panic(err)
	}
	return s
}

func (s *grpcServer) Run() errors.Error {
	var sererOptions []grpc.ServerOption
	s.gRpcServer = grpc.NewServer(sererOptions...)
	auth.RegisterAuthServer(s.gRpcServer, s)
	l, err := net.Listen("tcp", s.config.ListenAddr)
	if err != nil {
		return errors.New(err)
	}
	s.logger.Infof("grpc server listening to %s", s.config.ListenAddr)
	if err = s.gRpcServer.Serve(l); err != nil {
		return errors.New(err)
	}
	return nil
}

func (s *grpcServer) Stop() {
	s.gRpcServer.Stop()
}
