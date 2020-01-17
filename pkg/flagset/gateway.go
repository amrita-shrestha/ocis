package flagset

import (
	"github.com/micro/cli"
	"github.com/owncloud/ocis-reva/pkg/config"
)

// GatewayWithConfig applies cfg to the root flagset
func GatewayWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{

		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Usage:       "Enable sending traces",
			EnvVar:      "REVA_TRACING_ENABLED",
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-type",
			Value:       "jaeger",
			Usage:       "Tracing backend type",
			EnvVar:      "REVA_TRACING_TYPE",
			Destination: &cfg.Tracing.Type,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "Endpoint for the agent",
			EnvVar:      "REVA_TRACING_ENDPOINT",
			Destination: &cfg.Tracing.Endpoint,
		},
		&cli.StringFlag{
			Name:        "tracing-collector",
			Value:       "",
			Usage:       "Endpoint for the collector",
			EnvVar:      "REVA_TRACING_COLLECTOR",
			Destination: &cfg.Tracing.Collector,
		},
		&cli.StringFlag{
			Name:        "tracing-service",
			Value:       "reva",
			Usage:       "Service name for tracing",
			EnvVar:      "REVA_TRACING_SERVICE",
			Destination: &cfg.Tracing.Service,
		},

		// debug ports are the odd ports
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9143",
			Usage:       "Address to bind debug server",
			EnvVar:      "REVA_GATEWAY_DEBUG_ADDR",
			Destination: &cfg.Reva.Gateway.DebugAddr,
		},
		&cli.StringFlag{
			Name:        "debug-token",
			Value:       "",
			Usage:       "Token to grant metrics access",
			EnvVar:      "REVA_DEBUG_TOKEN",
			Destination: &cfg.Debug.Token,
		},
		&cli.BoolFlag{
			Name:        "debug-pprof",
			Usage:       "Enable pprof debugging",
			EnvVar:      "REVA_DEBUG_PPROF",
			Destination: &cfg.Debug.Pprof,
		},
		&cli.BoolFlag{
			Name:        "debug-zpages",
			Usage:       "Enable zpages debugging",
			EnvVar:      "REVA_DEBUG_ZPAGES",
			Destination: &cfg.Debug.Zpages,
		},

		// REVA

		&cli.StringFlag{
			Name:        "jwt-secret",
			Value:       "Pive-Fumkiu4",
			Usage:       "Shared jwt secret for reva service communication",
			EnvVar:      "REVA_JWT_SECRET",
			Destination: &cfg.Reva.JWTSecret,
		},
		&cli.StringFlag{
			Name:        "transfer-secret",
			Value:       "replace-me-with-a-transfer-secret",
			Usage:       "Transfer secret for datagateway",
			EnvVar:      "REVA_TRANSFER_SECRET",
			Destination: &cfg.Reva.TransferSecret,
		},
		&cli.IntFlag{
			Name:        "transfer-expires",
			Value:       10,
			Usage:       "Transfer secret for datagateway",
			EnvVar:      "REVA_TRANSFER_EXPIRES",
			Destination: &cfg.Reva.TransferExpires,
		},

		// TODO allow configuring clients

		// Services

		// Gateway

		&cli.StringFlag{
			Name:        "network",
			Value:       "tcp",
			Usage:       "Network to use for the reva service, can be 'tcp', 'udp' or 'unix'",
			EnvVar:      "REVA_GATEWAY_NETWORK",
			Destination: &cfg.Reva.Gateway.Network,
		},
		&cli.StringFlag{
			Name:        "protocol",
			Value:       "grpc",
			Usage:       "protocol for reva service, can be 'http' or 'grpc'",
			EnvVar:      "REVA_GATEWAY_PROTOCOL",
			Destination: &cfg.Reva.Gateway.Protocol,
		},
		&cli.StringFlag{
			Name:        "addr",
			Value:       "0.0.0.0:9142",
			Usage:       "Address to bind reva service",
			EnvVar:      "REVA_GATEWAY_ADDR",
			Destination: &cfg.Reva.Gateway.Addr,
		},
		&cli.StringFlag{
			Name:        "url",
			Value:       "localhost:9142",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_GATEWAY_URL",
			Destination: &cfg.Reva.Gateway.URL,
		},
		&cli.StringFlag{
			Name:        "services",
			Value:       "gateway,authregistry,storageregistry", // TODO appregistry
			Usage:       "comma separated list of services to include",
			EnvVar:      "REVA_GATEWAY_SERVICES",
			Destination: &cfg.Reva.Gateway.Services,
		},
		// TODO should defaults to true. reverse logic to 'disable-share-commit'?
		&cli.BoolFlag{
			Name:        "commit-share-to-storage-grant",
			Usage:       "Commit shares to the share manager as well as as a grant to the storage",
			EnvVar:      "REVA_GATEWAY_COMMIT_SHARE_TO_STRORAGE_GRANT",
			Destination: &cfg.Reva.Gateway.CommitShareToStorageGrant,
		},

		// other services

		// storage registry

		&cli.StringFlag{
			Name:        "frontend-url",
			Value:       "localhost:9140",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_FRONTEND_URL",
			Destination: &cfg.Reva.Frontend.URL,
		},
		&cli.StringFlag{
			Name:        "users-url",
			Value:       "localhost:9144",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_USERS_URL",
			Destination: &cfg.Reva.Users.URL,
		},
		&cli.StringFlag{
			Name:        "auth-basic-url",
			Value:       "localhost:9146",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_AUTH_BASIC_URL",
			Destination: &cfg.Reva.AuthBasic.URL,
		},
		&cli.StringFlag{
			Name:        "auth-bearer-url",
			Value:       "localhost:9148",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_AUTH_BEARER_URL",
			Destination: &cfg.Reva.AuthBearer.URL,
		},
		&cli.StringFlag{
			Name:        "sharing-url",
			Value:       "localhost:9150",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_SHARING_URL",
			Destination: &cfg.Reva.Sharing.URL,
		},

		&cli.StringFlag{
			Name:        "storage-root-url",
			Value:       "localhost:9152",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_STORAGE_ROOT_URL",
			Destination: &cfg.Reva.StorageRoot.URL,
		},
		&cli.StringFlag{
			Name:        "storage-root-mount-path",
			Value:       "/",
			Usage:       "mount path",
			EnvVar:      "REVA_STORAGE_ROOT_MOUNT_PATH",
			Destination: &cfg.Reva.StorageRoot.MountPath,
		},
		&cli.StringFlag{
			Name:        "storage-root-mount-id",
			Value:       "123e4567-e89b-12d3-a456-426655440001",
			Usage:       "mount id",
			EnvVar:      "REVA_STORAGE_ROOT_MOUNT_ID",
			Destination: &cfg.Reva.StorageRoot.MountID,
		},

		&cli.StringFlag{
			Name:        "storage-home-url",
			Value:       "localhost:9154",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_STORAGE_HOME_URL",
			Destination: &cfg.Reva.StorageHome.URL,
		},
		&cli.StringFlag{
			Name:        "storage-home-mount-path",
			Value:       "/home",
			Usage:       "mount path",
			EnvVar:      "REVA_STORAGE_HOME_MOUNT_PATH",
			Destination: &cfg.Reva.StorageHome.MountPath,
		},
		&cli.StringFlag{
			Name:        "storage-home-mount-id",
			Value:       "123e4567-e89b-12d3-a456-426655440000",
			Usage:       "mount id",
			EnvVar:      "REVA_STORAGE_HOME_MOUNT_ID",
			Destination: &cfg.Reva.StorageHome.MountID,
		},

		&cli.StringFlag{
			Name:        "storage-home-data-url",
			Value:       "localhost:9156",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_STORAGE_HOME_DATA_URL",
			Destination: &cfg.Reva.StorageHomeData.URL,
		},

		&cli.StringFlag{
			Name:        "storage-oc-url",
			Value:       "localhost:9162",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_STORAGE_OC_URL",
			Destination: &cfg.Reva.StorageOC.URL,
		},
		&cli.StringFlag{
			Name:        "storage-oc-mount-path",
			Value:       "/oc",
			Usage:       "mount path",
			EnvVar:      "REVA_STORAGE_OC_MOUNT_PATH",
			Destination: &cfg.Reva.StorageOC.MountPath,
		},
		&cli.StringFlag{
			Name:        "storage-oc-mount-id",
			Value:       "123e4567-e89b-12d3-a456-426655440000",
			Usage:       "mount id",
			EnvVar:      "REVA_STORAGE_OC_MOUNT_ID",
			Destination: &cfg.Reva.StorageOC.MountID,
		},

		&cli.StringFlag{
			Name:        "storage-oc-data-url",
			Value:       "localhost:9164",
			Usage:       "URL to use for the reva service",
			EnvVar:      "REVA_STORAGE_OC_DATA_URL",
			Destination: &cfg.Reva.StorageOCData.URL,
		},
	}
}
