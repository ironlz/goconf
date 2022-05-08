package goconf

import (
	"fmt"
	"testing"
)

func TestConstructConfig(t *testing.T) {
	config, err := ConstructConfig("../configs")
	if err != nil {
		t.Fatalf("Failed parse config, %v", err)
	}

	fmt.Println(config)
}
