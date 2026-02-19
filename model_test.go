package main


import (
	"testing"
)

func TestValidateTodo(t *testing.T) {
    tests := []struct {
        name string
        todo Todo
        wantErr bool
    }{
        {"valid", Todo{Title: "Task", Importance: 3}, false},
        {"empty title", Todo{Title: "", Importance: 3}, true},
        {"importance low", Todo{Title: "X", Importance: 0}, true},
        {"importance high", Todo{Title: "X", Importance: 6}, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateTodo(&tt.todo)
            if (err != nil) != tt.wantErr {
                t.Errorf("expected error=%v got=%v", tt.wantErr, err)
            }
        })
    }
}

func TestToggleCompleted(t *testing.T) {
    todo := Todo{Completed: false}
    todo.ToggleCompleted()
    if !todo.Completed {
        t.Errorf("expected true after toggle")
    }

    todo.ToggleCompleted()
    if todo.Completed {
        t.Errorf("expected false after second toggle")
    }
}

func TestValidateTodoImportanceLow(t *testing.T) {
    todo := Todo{Title: "X", Importance: 0}
    if err := validateTodo(&todo); err == nil {
        t.Errorf("expected error for low importance")
    }
}

func TestValidateTodoImportanceHigh(t *testing.T) {
    todo := Todo{Title: "X", Importance: 10}
    if err := validateTodo(&todo); err == nil {
        t.Errorf("expected error for high importance")
    }
}

func TestValidateTodoWhitespaceTitle(t *testing.T) {
    todo := Todo{Title: "   ", Importance: 3}
    if err := validateTodo(&todo); err == nil {
        t.Errorf("expected error for whitespace title")
    }
}




