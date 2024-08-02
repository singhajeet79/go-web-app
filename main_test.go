package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// Mocking the renderTemplate function for testing
func renderTemplate(w http.ResponseWriter, tmpl string, title string) {
    // This is a simplified version just for testing
    // You can add a simple HTML string here for testing
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    
    // Check the error returned by w.Write
    _, err := w.Write([]byte("<html><head><title>" + title + "</title></head><body>Home Page</body></html>"))
    if err != nil {
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
    }
}

func TestHome(t *testing.T) {
    req, err := http.NewRequest("GET", "/home", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "home.html", "Welcome to DevOps Tools 2024")
    })

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := "text/html; charset=utf-8"
    if contentType := rr.Header().Get("Content-Type"); contentType != expected {
        t.Errorf("handler returned unexpected content type: got %v want %v",
            contentType, expected)
    }
}
