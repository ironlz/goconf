package conf

import (
	"fmt"
	"testing"
)

func TestConstructConfig(t *testing.T) {
	config, err := ConstructConfig("../configs")
	if err != nil {
		t.Fatalf("Failed parse config, %v", err)
	}

	property := config.Use("config").GetProperty("app.listen")
	fmt.Println(property)
}
