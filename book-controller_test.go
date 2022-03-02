package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raghavi101/go-store/pkg/controllers"
)

func TestGetBook(t *testing.T) {
	req, err := http.NewRequest("GET", "/book/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetBook)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Welcome to crud api[{"ID":2,"CreatedAt":"2022-02-16T18:26:43+05:30","UpdatedAt":"2022-02-16T18:26:43+05:30","DeletedAt":null,"name":"abc","author":"abc","publication":"abc"},{"ID":3,"CreatedAt":"2022-02-24T17:22:32+05:30","UpdatedAt":"2022-02-24T17:22:32+05:30","DeletedAt":null,"name":"xyzabc","author":"xyzabc","publication":"xyzabc"},{"ID":4,"CreatedAt":"2022-03-02T11:54:02+05:30","UpdatedAt":"2022-03-02T11:54:02+05:30","DeletedAt":null,"name":"qwerty1","author":"qwerty1","publication":"qwerty1"},{"ID":5,"CreatedAt":"2022-03-02T11:55:08+05:30","UpdatedAt":"2022-03-02T11:55:08+05:30","DeletedAt":null,"name":"qwerty1","author":"qwerty1","publication":"qwerty1"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetBookById(t *testing.T) {

	req, err := http.NewRequest("GET", "/book/{bookId}", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "2")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetBookById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"","author":"","publication":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetBookByIdNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/book/{bookId}", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "900")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetBookById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCreateBook(t *testing.T) {

	var jsonStr = []byte(`{"name":"qwerty1","author":"qwerty1","publication":"qwerty1"}`)

	req, err := http.NewRequest("POST", "/book/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "pkglication/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateBook)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":6,"CreatedAt":"2022-03-02T12:09:03.824468+05:30","UpdatedAt":"2022-03-02T12:09:03.824468+05:30","DeletedAt":null,"name":"qwerty1","author":"qwerty1","publication":"qwerty1"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDeleteBook(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/book/{bookId}", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "3")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.DeleteBook)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"","author":"","publication":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateBook(t *testing.T) {

	var jsonStr = []byte(`{"name":"qwerty1","author":"qwerty1change","publication":"qwerty1"}`)

	req, err := http.NewRequest("PUT", "/book/{bookId}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.UpdateBook)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"qwerty1","author":"qwerty1change","publication":"qwerty1"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
