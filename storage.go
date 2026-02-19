package main

type Storage interface {
    Add(Todo) Todo
    List() []Todo
    Get(id int) (Todo, bool)
}

type MemoryStorage struct{}

func (MemoryStorage) Add(t Todo) Todo {
    t.ID = nextID
    nextID++
    todos = append(todos, t)
    return t
}

func (MemoryStorage) List() []Todo {
    return todos
}

func (MemoryStorage) Get(id int) (Todo, bool) {
    for _, t := range todos {
        if t.ID == id {
            return t, true
        }
    }
    return Todo{}, false
}