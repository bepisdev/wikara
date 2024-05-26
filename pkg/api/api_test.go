package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeHandler_ValidPath(t *testing.T) {
	handlerFn := func(w http.ResponseWriter, r *http.Request, id string) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ID: " + id))
	}

	// Create a request with a valid path
	req, err := http.NewRequest("GET", "/edit/test123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := MakeHandler(handlerFn)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "ID: test123"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMakeHandler_InvalidPath(t *testing.T) {
	handlerFn := func(w http.ResponseWriter, r *http.Request, id string) {
		t.Error("Handler should not be called for invalid paths")
	}

	// Create a request with an invalid path
	req, err := http.NewRequest("GET", "/invalid/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := MakeHandler(handlerFn)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
