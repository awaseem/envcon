package main

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	if err := os.Setenv("gettest", "this is a test"); err != nil {
		t.Error("Failed to set env for testing")
	}

	env := &envStorage{}

	str, err := env.get("gettest")
	if err != nil {
		t.Error("env get method threw an error when it was not suppose too!")
	}
	if str != "this is a test" {
		t.Errorf("env get method got different test")
	}
}

func TestSetEnv(t *testing.T) {
	env := &envStorage{}

	err := env.set("settest", "this is a test")
	if err != nil {
		t.Error("env get method threw an error when it was not suppose too!")
	}
	if os.Getenv("settest") != "this is a test" {
		t.Errorf("env get method got different test")
	}
}
