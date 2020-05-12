package config

import (
	"google.golang.org/grpc"

	"proxy-server/internal/proto/auth"
)

var (
	AuthClient auth.AuthClient
)

func ConnectAuthService() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		SystemLogger.Error(err)
	}

	AuthClient = auth.NewAuthClient(conn)
}
