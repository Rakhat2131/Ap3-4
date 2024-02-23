package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/usecase"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	host := "localhost"
	port := "5433"
	user := "postgres"
	password := "123456"
	dbname := "postgres"

	db, err := postgres.ConnectDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	defer db.Close()

	fmt.Println("Connected to the database!")

	contactRepo := repository.NewRepository()
	contactUC := usecase.NewContactUseCase()
	contactHandler := delivery.NewDelivery()

	fmt.Println(contactRepo)
	fmt.Println(contactUC)
	fmt.Println(contactHandler)

	startServer(contactHandler)

	// Start your server, passing in the handler
	startServer(contactHandler)

}

func startServer(handler *delivery.Delivery) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
