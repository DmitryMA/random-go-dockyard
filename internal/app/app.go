package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DmitryMA/go-Chi-PostgreSQL-Docker/internal/api"
	"github.com/DmitryMA/go-Chi-PostgreSQL-Docker/internal/store"
	"github.com/DmitryMA/go-Chi-PostgreSQL-Docker/migrations"
)

type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB *sql.DB
}


func NewApplication() (*Application, error) {
	pgDB, err := store.Open()

	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")

	if  err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
 
	// our stores
	
	// our handlers
	workoutHandler := api.NewWorkoutHandler()

	app := &Application {	
 	 	Logger: logger,
		WorkoutHandler: workoutHandler,
		DB: pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
