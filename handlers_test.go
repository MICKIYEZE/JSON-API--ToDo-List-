package main

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandleAdd(t *testing.T) {
    body := `{"title":"Test","importance":3}`
    req := httptest.NewRequest("POST", "/add", bytes.NewBufferString(body))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    handleAdd(w, req)

    if w.Code != http.StatusCreated {
        t.Errorf("expected 201 got %d", w.Code)
    }
}

func TestHandleAddInvalidJSON(t *testing.T) {
    req := httptest.NewRequest("POST", "/add", bytes.NewBufferString("{bad json"))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    handleAdd(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("expected 400 got %d", w.Code)
    }
}

func TestHandleItemNotFound(t *testing.T) {
    todos = []Todo{}
    req := httptest.NewRequest("GET", "/item/99", nil)
    w := httptest.NewRecorder()

    handleItem(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("expected 404 got %d", w.Code)
    }
}

func TestHandleAddWrongMethod(t *testing.T) {
    req := httptest.NewRequest("GET", "/add", nil)
    w := httptest.NewRecorder()

    handleAdd(w, req)

    if w.Code != http.StatusMethodNotAllowed {
        t.Errorf("expected 405, got %d", w.Code)
    }
}

func TestHandleItemInvalidID(t *testing.T) {
    req := httptest.NewRequest("GET", "/item/abc", nil)
    w := httptest.NewRecorder()

    handleItem(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("expected 400, got %d", w.Code)
    }
}

func TestHandleItemWrongMethod(t *testing.T) {
    req := httptest.NewRequest("POST", "/item/1", nil)
    w := httptest.NewRecorder()

    handleItem(w, req)

    if w.Code != http.StatusMethodNotAllowed {
        t.Errorf("expected 405, got %d", w.Code)
    }
}

func TestHandleAddMissingContentType(t *testing.T) {
    body := `{"title":"X","importance":3}`
    req := httptest.NewRequest("POST", "/add", bytes.NewBufferString(body))
    w := httptest.NewRecorder()

    handleAdd(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("expected 400, got %d", w.Code)
    }
}

func TestHandleItemInvalidPath(t *testing.T) {
    req := httptest.NewRequest("GET", "/item", nil)
    w := httptest.NewRecorder()

    handleItem(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("expected 404, got %d", w.Code)
    }
}

func TestHandleListWrongMethod(t *testing.T) {
    req := httptest.NewRequest("POST", "/list", nil)
    w := httptest.NewRecorder()

    handleList(w, req)

    if w.Code != http.StatusMethodNotAllowed {
        t.Errorf("expected 405, got %d", w.Code)
    }
}

func TestHandleRootWrongMethod(t *testing.T) {
    req := httptest.NewRequest("POST", "/", nil)
    w := httptest.NewRecorder()

    handleRoot(w, req)

    if w.Code != http.StatusMethodNotAllowed {
        t.Errorf("expected 405, got %d", w.Code)
    }
}

func TestHandleItemDoubleSlash(t *testing.T) {
    req := httptest.NewRequest("GET", "/item//5", nil)
    w := httptest.NewRecorder()

    handleItem(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("expected 404, got %d", w.Code)
    }
}

func TestHandleAddValidButFailsValidation(t *testing.T) {
    body := `{"title":"", "importance":3}`
    req := httptest.NewRequest("POST", "/add", bytes.NewBufferString(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    handleAdd(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("expected 400, got %d", w.Code)
    }
}

func TestHandleItemTrailingSlash(t *testing.T) {
    req := httptest.NewRequest("GET", "/item/", nil)
    w := httptest.NewRecorder()

    handleItem(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("expected 404, got %d", w.Code)
    }
}

func TestHandleItemTooManySegments(t *testing.T) {
    req := httptest.NewRequest("GET", "/item/1/extra", nil)
    w := httptest.NewRecorder()

    handleItem(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("expected 404, got %d", w.Code)
    }
}




