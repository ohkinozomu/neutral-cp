package neutral_cp

import (
	"context"

	"cloud.google.com/go/profiler"
)

func (n *NeutralCP) startCloudProfiler(ctx context.Context) error {
	cfg := profiler.Config{
		Service:        n.Config.ApplicationName,
		ServiceVersion: n.Config.Version,
	}

	if err := profiler.Start(cfg); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
