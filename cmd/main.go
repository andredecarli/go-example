package main

import (
	"log"
	"net/http"

	customerService "github.com/andredecarli/go-example/internal/application/customer"
	customerController "github.com/andredecarli/go-example/internal/controllers/http/customer"
	"github.com/andredecarli/go-example/internal/infra/config"
	customerRepository "github.com/andredecarli/go-example/internal/infra/db/mongodb/customer"
)

func main() {
	log.Println("System started.")

	cfg := config.LoadConfig()

	client, err := config.NewMongoClient(cfg)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	db := client.Database(cfg.MongoDatabase)

	log.Println("MongoDB client created.")
	customerRepository := customerRepository.NewRepository(db)
	log.Println("Customer repository created.")

	customerService := customerService.NewService(customerRepository)
	log.Println("Customer service created.")

	customerHandler := customerController.NewHandler(customerService)
	log.Println("Customer handler created.")

	mux := http.NewServeMux()
	customerController.RegisterRoutes(mux, customerHandler)
	log.Println("Customer routes registered.")

	log.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
