package main

import (
	"context"
	"testing"
)

func TestLoginQuery(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Errorf("NewClient(): %s", err)
		t.FailNow()
	}

	var query struct {
		Viewer struct {
			Login string
		}
	}

	err = client.Query(context.Background(), &query, nil)
	if err != nil {
		t.Errorf("query: %s", err)
		t.FailNow()
	}
}
