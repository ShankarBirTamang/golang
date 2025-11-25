package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 1. extract slug from URL
	paramWorkoutID := chi.URLParam(r, "id")

	// 2. validate
	if paramWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	// 3. convert to int
	workoutID, err := strconv.ParseInt(paramWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// 4. respond with message
	fmt.Fprintf(w, "This is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintf(w, "Created a workout\n")
}
