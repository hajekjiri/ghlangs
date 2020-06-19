package main

import (
	"testing"
)

func TestCreateClient(t *testing.T) {
	_, err := NewClient()
	if err != nil {
		t.Errorf("NewClient(): %s", err)
		t.Fail()
	}
}
