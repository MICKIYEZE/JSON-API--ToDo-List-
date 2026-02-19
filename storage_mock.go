package main

type MockStorage struct {
    AddFunc  func(Todo) Todo
    ListFunc func() []Todo
    GetFunc  func(int) (Todo, bool)
}

func (m MockStorage) Add(t Todo) Todo {
    return m.AddFunc(t)
}

func (m MockStorage) List() []Todo {
    return m.ListFunc()
}

func (m MockStorage) Get(id int) (Todo, bool) {
    return m.GetFunc(id)
}