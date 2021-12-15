package grpc

import (
	accountssvc "github.com/owncloud/ocis/protogen/gen/ocis/services/accounts/v1"

	"github.com/owncloud/ocis/ocis-pkg/service/grpc"
)

// Server initializes a new go-micro service ready to run
func Server(opts ...Option) grpc.Service {
	options := newOptions(opts...)
	handler := options.Handler

	service := grpc.NewService(
		grpc.Name(options.Config.Server.Name),
		grpc.Context(options.Context),
		grpc.Address(options.Config.GRPC.Addr),
		grpc.Namespace(options.Config.GRPC.Namespace),
		grpc.Logger(options.Logger),
		grpc.Flags(options.Flags...),
		grpc.Version(options.Config.Server.Version),
	)

	if err := accountssvc.RegisterAccountsServiceHandler(service.Server(), handler); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register service handler")
	}
	if err := accountssvc.RegisterGroupsServiceHandler(service.Server(), handler); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register groups handler")
	}
	if err := accountssvc.RegisterIndexServiceHandler(service.Server(), handler); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register index handler")
	}

	return service
}
