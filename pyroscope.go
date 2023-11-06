package neutral_cp

import (
	"context"

	"github.com/grafana/pyroscope-go"
)

func (n *NeutralCP) startPyroscope(ctx context.Context) error {
	p, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: n.Config.ApplicationName,

		ServerAddress: n.Config.ServerAddress,

		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
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
