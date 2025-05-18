package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/DmitryMA/go-Chi-PostgreSQL-Docker/internal/app"
	"github.com/DmitryMA/go-Chi-PostgreSQL-Docker/internal/routes"
)

func main()	{
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	application, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	defer application.DB.Close()
	
	r := routes.SetupRouts(application)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	
	application.Logger.Printf("We are running on port: %d our App \n", port)

	err = server.ListenAndServe()

	if err != nil {
		application.Logger.Fatal(err)
	}
}
