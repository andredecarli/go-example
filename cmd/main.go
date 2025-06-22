package main

import (
	"log"
	"time"

	"github.com/andredecarli/go-example/internal/application/services"
	"github.com/andredecarli/go-example/internal/controllers"
	"github.com/andredecarli/go-example/internal/infra/db/memory"
)

func main() {
	log.Println("System is running...")
	time.Sleep(1 * time.Second)

	customerRepository := memory.NewCustomerRepository()
	log.Println("Customer repository created")
	time.Sleep(1 * time.Second)

	customerService := services.NewCustomerService(customerRepository)
	log.Println("Customer service created")
	time.Sleep(1 * time.Second)

	customerController := controllers.NewCustomerController(customerService)
	log.Println("Customer controller created")
	time.Sleep(2 * time.Second)

	log.Println("Creating customer")
	created, err := customerController.Create("John Doe", "john@doe.com")
	if err != nil {
		log.Println("Error creating customer:", err)
	}
	log.Println("Customer created:", created)
	time.Sleep(5 * time.Second)

	log.Println("Listing customers")
	customers, err := customerController.List()
	if err != nil {
		log.Println("Error listing customers:", err)
	}
	log.Println("Customers:", customers)
	time.Sleep(5 * time.Second)

	log.Println("Finding customer by ID")
	found, err := customerController.FindByID(created.ID)
	if err != nil {
		log.Println("Error finding customer:", err)
	}
	log.Println("Customer found:", found)
	time.Sleep(5 * time.Second)

	log.Println("Updating customer")
	updated, err := customerController.Update(created.ID, "Jane Doe", "jane@doe.com")
	if err != nil {
		log.Println("Error updating customer:", err)
	}
	log.Println("Customer updated:", updated)
	time.Sleep(5 * time.Second)

	log.Println("Deleting customer")
	err = customerController.Delete(created.ID)
	if err != nil {
		log.Println("Error deleting customer:", err)
	}
	log.Println("Customer deleted")
	time.Sleep(5 * time.Second)

	log.Println("Finding customer by ID")
	found, err = customerController.FindByID(created.ID)
	if err != nil {
		log.Println("Error finding customer:", err)
	}
	log.Println("Customer found:", found)
	time.Sleep(5 * time.Second)

	log.Println("Listing customers")
	customers, err = customerController.List()
	if err != nil {
		log.Println("Error listing customers:", err)
	}
	log.Println("Customers:", customers)
	time.Sleep(2 * time.Second)

	log.Println("System shut down.")
}
