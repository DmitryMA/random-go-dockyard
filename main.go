package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DmitryMA/go-Chi-PostgreSQL-Docker/internal/app"
)

func main()	{
	application, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	application.Logger.Println("We are running our App")

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr: ":8080",
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		application.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}