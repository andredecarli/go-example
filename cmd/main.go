package main

import (
	"context"
	"log"
	"time"

	customerService "github.com/andredecarli/go-example/internal/application/customer"
	customerController "github.com/andredecarli/go-example/internal/controllers/core/customer"
	customerRepository "github.com/andredecarli/go-example/internal/infra/db/memory/customer"
)

func main() {
	log.Println("System started.")
	ctx := context.Background()
	time.Sleep(1 * time.Second)

	customerRepository := customerRepository.NewRepository()
	log.Println("Customer repository created.")
	time.Sleep(1 * time.Second)

	customerService := customerService.NewService(customerRepository)
	log.Println("Customer service created.")
	time.Sleep(1 * time.Second)

	customerController := customerController.NewCustomerController(customerService)
	log.Println("Customer controller created.")
	time.Sleep(2 * time.Second)

	customerController.Create(ctx, "John Doe", "john.doe@example.com")
	time.Sleep(5 * time.Second)

	log.Println("System shut down.")
}
