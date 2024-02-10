package main

import (
	"github.com/sdk-fabric/twitter-go/sdk"
	"testing"
)

func TestClient(t *testing.T) {

	client, err := sdk.Build("my_token")
	if err != nil {
		t.Error(err)
	}

	if client == nil {
		t.Error("Client not defined")
	}

}
