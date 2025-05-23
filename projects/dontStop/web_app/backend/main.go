// backend/main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type ExerciseType string

const (
	Stretch ExerciseType = "stretch"
	General ExerciseType = "general"
)

type Exercise struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Type       ExerciseType `json:"type"`
	HoldSec    int          `json:"holdSec,omitempty"`    // for stretch
	RestSec    int          `json:"restSec,omitempty"`    // for stretch
	Limit      int          `json:"limit,omitempty"`      // for stretch (cycles)
	RepsPerSet int          `json:"repsPerSet,omitempty"` // for general
	TotalSets  int          `json:"totalSets,omitempty"`  // for general
}

var (
	exercises = []Exercise{}
	nextID    = 1
	mu        sync.Mutex
)

func listExercises(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}

func addExercise(w http.ResponseWriter, r *http.Request) {
	var ex Exercise
	if err := json.NewDecoder(r.Body).Decode(&ex); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	ex.ID = nextID
	nextID++
	exercises = append(exercises, ex)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ex)
}

func main() {
	http.HandleFunc("/exercises", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			listExercises(w, r)
		} else if r.Method == http.MethodPost {
			addExercise(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
