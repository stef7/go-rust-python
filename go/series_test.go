package main

import (
	"fmt"
	"testing"
)

func TestGetUsers(t *testing.T) {
	expected := "[]main.User{main.User{Name:\"Alice\", Age:34, Active:true, LastLogin:time.Date(2023, time.December, 1, 10, 30, 0, 0, time.UTC), Children:[]main.User{main.User{Name:\"Bob\", Age:29, Active:true, LastLogin:time.Date(2023, time.November, 15, 8, 15, 0, 0, time.UTC), Children:[]main.User{}}, main.User{Name:\"Carol\", Age:38, Active:false, LastLogin:time.Date(2023, time.October, 20, 18, 0, 0, 0, time.UTC), Children:[]main.User{}}}}, main.User{Name:\"Dave\", Age:45, Active:true, LastLogin:time.Date(2023, time.November, 20, 17, 45, 0, 0, time.UTC), Children:[]main.User{}}, main.User{Name:\"Eve\", Age:28, Active:true, LastLogin:time.Date(2024, time.January, 10, 9, 0, 0, 0, time.UTC), Children:[]main.User{main.User{Name:\"Frank\", Age:31, Active:true, LastLogin:time.Date(2024, time.January, 12, 11, 30, 0, 0, time.UTC), Children:[]main.User{main.User{Name:\"Grace\", Age:27, Active:false, LastLogin:time.Date(2023, time.September, 25, 16, 45, 0, 0, time.UTC), Children:[]main.User{}}}}}}, main.User{Name:\"Heidi\", Age:50, Active:false, LastLogin:time.Date(2022, time.June, 30, 12, 0, 0, 0, time.UTC), Children:[]main.User{}}}"

	users := GetUsers()
	returned := fmt.Sprintf("%#v", users)

	if returned != expected {
		t.Errorf("Not equal:\nExpected: %v\nReturned: %v", expected, returned)
	}
}
