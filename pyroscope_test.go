package neutral_cp

import (
	"context"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	c := Config{
		registry:        "pyroscope",
		applicationName: "testApp",
	}
	ncp := NeutralCP{config: c}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err := ncp.Start(ctx)
	if err != nil {
		t.Fatalf("failed test")
	}
}
