package util

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupMongoDBContainer() (context.Context, testcontainers.Container, string, error) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
	}
	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, "", err
	}

	port, err := mongoC.MappedPort(ctx, "27017")
	if err != nil {
		return nil, nil, "", err
	}

	uri := fmt.Sprintf("mongodb://localhost:%s", port.Port())
	return ctx, mongoC, uri, nil
}
