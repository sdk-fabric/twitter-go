package main

import (
	"github.com/sdk-fabric/twitter-go/sdk"
	"github.com/apioo/sdkgen-go/v2"
	"testing"
)

func TestClient(t *testing.T) {

	client, err := sdk.Build("my_token")
	if err != nil {
		t.Error(err)
	}

}
