package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {

	t.Run("Tabular tests", func(t *testing.T) {
		t.Parallel()
		for i := 1; i <= 5; i++ {
			t.Run(fmt.Sprintf("it returns a single user with id %d", i), func(t *testing.T) {
				req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", i), nil)
				res := httptest.NewRecorder()

				ServeHTTP(res, req)
				AssertEqual(t, res.Code, http.StatusOK)

				var got User
				AssertNil(t, json.NewDecoder(res.Body).Decode(&got))
				AssertEqual(t, got.ID, fmt.Sprintf("%d", i))
			})
		}
	})

	t.Run("it returns a list of users", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/users", nil)
		res := httptest.NewRecorder()

		ServeHTTP(res, req)
		AssertEqual(t, res.Code, http.StatusOK)

		var got []User
		AssertNil(t, json.NewDecoder(res.Body).Decode(&got))

		if len(got) == 0 {
			t.Errorf("Expected atleast 1 user but %d", len(got))
		}
	})
}

func TestCreateUser(t *testing.T) {

	var id string
	t.Run("it returns accepted for creating a user", func(t *testing.T) {

		user := User{Name: "Olu"}
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&user)

		req := httptest.NewRequest("POST", "/users", buffer)
		res := httptest.NewRecorder()

		ServeHTTP(res, req)
		AssertEqual(t, res.Code, http.StatusAccepted)

		var got User
		AssertNil(t, json.NewDecoder(res.Body).Decode(&got))
		AssertEqual(t, got.Name, user.Name)
		id = got.ID
	})

	t.Run("it returns OK for updating a user", func(t *testing.T) {

		user := User{Name: "Peter"}
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&user)

		req := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", id), buffer)
		res := httptest.NewRecorder()

		ServeHTTP(res, req)
		AssertEqual(t, res.Code, http.StatusOK)

		var got User
		AssertNil(t, json.NewDecoder(res.Body).Decode(&got))
		AssertEqual(t, got.Name, user.Name)
	})
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func AssertNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected nil error but got %v", err)
	}
}
