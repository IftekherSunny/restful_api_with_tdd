package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iftekhersunny/restful_api_with_tdd/routes"
	"github.com/stretchr/testify/assert"
)

////////////////////////////////////////////
// Dummy data
////////////////////////////////////////////
var defaultTodos = map[string]interface{}{
	"1": "Task one",
	"2": "Task two",
}

var defaultSingleTodo = "Task one"

////////////////////////////////////////////
// Api routes
////////////////////////////////////////////
var router = routes.Api()

func TestRestfulApiApp(t *testing.T) {

	t.Run("Test API Endpoint [GET]:/v1/todos", func(t *testing.T) {
		testTodosGetApi(t)
	})

	t.Run("Test API Endpoint [POST]:/v1/todos", func(t *testing.T) {
		testTodoCreateApi(t)
	})

	t.Run("Test API Endpoint [GET]:/v1/todos/1", func(t *testing.T) {
		testTodoGetApi(t)
	})

	t.Run("Test API Endpoint [PUT]:/v1/todos/1", func(t *testing.T) {
		testTodoUpdateApi(t)
	})

	t.Run("Test API Endpoint [DELETE]:/v1/todos/1", func(t *testing.T) {
		testTodoDeleteApi(t)
	})
}

func testTodosGetApi(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/v1/todos", nil)
	router.ServeHTTP(w, r)

	response := map[string]interface{}{}

	json.NewDecoder(w.Body).Decode(&response)

	if response["status"].(float64) != http.StatusOK {
		t.Error("Status code should be 200")
	}

	todos := response["data"].(map[string]interface{})["todos"]

	assert.Equal(t, todos, defaultTodos)
}

func testTodoCreateApi(t *testing.T) {
	testTodo := map[string]interface{}{"name": "Test todo"}
	data, _ := json.Marshal(testTodo)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/v1/todos", bytes.NewBufferString(string(data)))
	router.ServeHTTP(w, r)

	response := map[string]interface{}{}

	json.NewDecoder(w.Body).Decode(&response)

	if response["status"].(float64) != http.StatusCreated {
		t.Error("Status code should be 201")
	}

	assert.Equal(t, "New todo has been added", response["message"])

	assert.Equal(t, float64(3), response["total_todos"])
}

func testTodoGetApi(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/v1/todos/1", nil)
	router.ServeHTTP(w, r)

	response := map[string]interface{}{}

	json.NewDecoder(w.Body).Decode(&response)

	if response["status"].(float64) != http.StatusOK {
		t.Error("Status code should be 200")
	}

	todo := response["data"].(map[string]interface{})["todo"]

	assert.Equal(t, todo, defaultSingleTodo)
}

func testTodoUpdateApi(t *testing.T) {
	testTodo := map[string]interface{}{"name": "Test todo"}
	data, _ := json.Marshal(testTodo)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/v1/todos/1", bytes.NewBufferString(string(data)))
	router.ServeHTTP(w, r)

	response := map[string]interface{}{}

	json.NewDecoder(w.Body).Decode(&response)

	if response["status"].(float64) != http.StatusOK {
		t.Error("Status code should be 200")
	}

	assert.Equal(t, "New todo has been updated", response["message"])
}

func testTodoDeleteApi(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/v1/todos/1", nil)
	router.ServeHTTP(w, r)

	response := map[string]interface{}{}

	json.NewDecoder(w.Body).Decode(&response)

	if response["status"].(float64) != http.StatusOK {
		t.Error("Status code should be 200")
	}

	assert.Equal(t, "New todo has been deleted", response["message"])

	assert.Equal(t, float64(2), response["total_todos"])
}
