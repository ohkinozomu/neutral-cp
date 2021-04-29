package neutral_cp

import (
	"context"

	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
)

func (n *NeutralCP) startPyroscope(ctx context.Context) error {
	p, err := profiler.Start(profiler.Config{
		ApplicationName: n.config.applicationName,

		ServerAddress: n.config.serverAddress,

		ProfileTypes: []profiler.ProfileType{
			profiler.ProfileCPU,
			profiler.ProfileAllocObjects,
			profiler.ProfileAllocSpace,
			profiler.ProfileInuseObjects,
			profiler.ProfileInuseSpace,
		},
	})
	if err != nil {
		return err
	}

	<-ctx.Done()
	if err := p.Stop(); err != nil {
		return err
	}
	return nil
}
