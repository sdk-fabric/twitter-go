package main

import (
	"github.com/sdk-fabric/twitter-go/sdk"
	"github.com/apioo/sdkgen-go"
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {

	var credentials = sdkgen.HttpBearer{
		Token: "",
	}

	client, err := sdk.NewClient("https://api.acme.com", credentials)
	if err != nil {
		t.Error(err)
	}

}
