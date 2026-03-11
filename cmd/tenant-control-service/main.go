
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/georgirtodorov/tenant-control-service/api"
	"github.com/georgirtodorov/tenant-control-service/internal/gprcserver"
	"github.com/georgirtodorov/tenant-control-service/internal/registry"
)

func main() {
	port := flag.String("p", "8080", "Port for the gRPC server")
	flag.Parse()

	fmt.Println("Tenant Control Service Initialized")

	// create a new gRPC server
	grpcServer := gprcserver.New(*port)

	// create a new registry service
	registryService := registry.New()

	// register the registry service with the gRPC server
	api.RegisterRegistryServer(grpcServer.GetServer(), registryService)

	// start the gRPC server
	if err := grpcServer.Start(); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
