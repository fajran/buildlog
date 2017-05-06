package buildlog

import (
	"testing"
)

func TestDbUri_Minimum(t *testing.T) {
	db := DbConfig{
		Host: "host",
		Name: "name",
	}

	uri := db.Uri()
	if uri != "postgres://host/name" {
		t.Fail()
	}
}

func TestDbUri_WithUsername(t *testing.T) {
	db := DbConfig{
		Host:     "host",
		Name:     "name",
		Username: "user",
	}

	uri := db.Uri()
	if uri != "postgres://user@host/name" {
		t.Fail()
	}
}
func TestDbUri_WithUsernameAndPassword(t *testing.T) {
	db := DbConfig{
		Host:     "host",
		Name:     "name",
		Username: "user",
		Password: "pass",
	}

	uri := db.Uri()
	if uri != "postgres://user:pass@host/name" {
		t.Fail()
	}
}
func TestDbUri_WithPassword(t *testing.T) {
	db := DbConfig{
		Host:     "host",
		Name:     "name",
		Password: "pass",
	}

	uri := db.Uri()
	if uri != "postgres://host/name" {
		t.Fail()
	}
}
func TestDbUri_WithOptions(t *testing.T) {
	db := DbConfig{
		Host:    "host",
		Name:    "name",
		Options: "options",
	}

	uri := db.Uri()
	if uri != "postgres://host/name?options" {
		t.Fail()
	}
}
func TestDbUri_WithUsernameAndPasswordAndOptions(t *testing.T) {
	db := DbConfig{
		Host:     "host",
		Name:     "name",
		Username: "user",
		Password: "pass",
		Options:  "options",
	}

	uri := db.Uri()
	if uri != "postgres://user:pass@host/name?options" {
		t.Fail()
	}
}
