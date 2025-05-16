package routes

import (
	"github.com/DmitryMA/go-Chi-PostgreSQL-Docker/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRouts(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)
	
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	return r
}
