package main

import "testing"

func TestMemoryStorageCRUD(t *testing.T) {
    todos = []Todo{}
    nextID = 1
    s := MemoryStorage{}

    // Add
    added := s.Add(Todo{Title: "A", Importance: 3})
    if added.ID != 1 {
        t.Errorf("expected ID=1 got %d", added.ID)
    }

    // List
    list := s.List()
    if len(list) != 1 {
        t.Errorf("expected 1 item got %d", len(list))
    }

    // Get
    got, ok := s.Get(1)
    if !ok || got.Title != "A" {
        t.Errorf("expected to find item")
    }

    _, ok = s.Get(999)
    if ok {
        t.Errorf("expected not found")
    }
}

func TestMemoryStorageIDIncrement(t *testing.T) {
    todos = []Todo{}
    nextID = 1
    s := MemoryStorage{}

    a := s.Add(Todo{Title: "A", Importance: 3})
    b := s.Add(Todo{Title: "B", Importance: 4})

    if a.ID != 1 || b.ID != 2 {
        t.Errorf("expected IDs 1 and 2, got %d and %d", a.ID, b.ID)
    }
}

func TestMemoryStorageGetMissing(t *testing.T) {
    todos = []Todo{}
    nextID = 1
    s := MemoryStorage{}

    _, ok := s.Get(999)
    if ok {
        t.Errorf("expected not found")
    }
}