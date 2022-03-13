package redis

import (
	"github.com/alicebob/miniredis"
	"log"
	"testing"
)

func TestInitiate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Init-Redis"},
	}

	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initiate(mr.Addr())
		})
	}
}
