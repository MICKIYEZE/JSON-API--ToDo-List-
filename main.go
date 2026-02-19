package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Neat template for a future task
type Todo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	Importance int    `json:"importance"`
}

// Memory storage
var (
	todos  = []Todo{}
	nextID = 1
)

// JSON writer
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// Error json
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// Input validator
func validateTodo(t *Todo) error {
	if strings.TrimSpace(t.Title) == "" {
		return errors.New("title is required")
	}
	if t.Importance < 1 || t.Importance > 5 {
		return errors.New("importance must be between 1 and 5")
	}
	return nil
}

// Lists your task using /list using method GET
func handleList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "use GET")
		return
	}
	writeJSON(w, http.StatusOK, todos)
}

// Flips the completed stat like a light switch
func (t *Todo) ToggleCompleted() {
    t.Completed = !t.Completed
}


// Adds a task to youtr todo using /add with method POST
func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "use POST")
		return
	}

	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		writeError(w, http.StatusBadRequest, "expected application/json")
		return
	}

	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if err := validateTodo(&t); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	t.ID = nextID
	nextID++
	todos = append(todos, t)

	writeJSON(w, http.StatusCreated, t)
}

// Search for a specific item using: /item/{id} with method GET
func handleItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "use GET")
		return
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 || parts[0] != "item" {
		writeError(w, http.StatusNotFound, "invalid path")
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	for _, t := range todos {
		if t.ID == id {
			writeJSON(w, http.StatusOK, t)
			return
		}
	}

	writeError(w, http.StatusNotFound, "task not found")
}

// GET /
func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "use GET")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"message": "JSON API by MICHAEL",
		"routes": []string{
			"GET  /list",
			"POST /add",
			"GET  /item/{id}",
		},
	})
}

func main() {
	// starter pack
	todos = append(todos, Todo{
		ID:         nextID,
		Title:      "Get this program to work",
		Completed:  true,
		Importance: 5,
	})
	nextID++

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/list", handleList)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/item", handleItem)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
