package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"api/handlers"
	"api/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/users", handlers.GetPeople)
	r.GET("/user/:id", handlers.GetPerson)

	return r
}

func TestGetUsers(t *testing.T) {
	r := setupRouter()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	resp := &models.UsersResponse{}
	people := &models.Person{Id: 1, Name: "John", Age: 31, City: "New York"}
	resp.Users = append(resp.Users, people)
	people = &models.Person{Id: 2, Name: "Doe", Age: 22, City: "Vancouver"}
	resp.Users = append(resp.Users, people)
	//expected := `[{"id":1,"name":"John","age":31,"city":"New York"},{"id":2,"name":"Doe","age":22,"city":"Vancouver"}]`
	got := &models.UsersResponse{}
	if err := proto.Unmarshal(recorder.Body.Bytes(), got); err != nil {
		t.Fatalf("failed to unmashal resp: %s", err)
	}
	if !proto.Equal(got, resp) {
		t.Errorf("handler returned unexpected body: got %v want %v", got, resp)
	}
}

func TestGetUserByID(t *testing.T) {
	r := setupRouter()
	id := "2"
	req, err := http.NewRequest("GET", "/user/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":2,"name":"Doe","age":22,"city":"Vancouver"}`
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	r := setupRouter()
	id := "3"
	req, err := http.NewRequest("GET", "/user/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	if status := recorder.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
