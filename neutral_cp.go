package neutral_cp

import (
	"context"
	"errors"
	"log"
)

type Config struct {
	Registry        Registry
	ApplicationName string
	ServerAddress   string
	Version         string
}

type NeutralCP struct {
	Config Config
}

func (n *NeutralCP) Start(ctx context.Context) error {
	log.Println("Start neutral-cp")

	switch n.Config.Registry {
	case PYROSCOPE:
		n.startPyroscope(ctx)
	case CLOUD_PROFILER:
		n.startCloudProfiler(ctx)
	default:
		return errors.New("unsupported registry")
	}

	return nil
}
