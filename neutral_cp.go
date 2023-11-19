package neutral_cp

import (
	"context"
	"errors"

	"go.uber.org/zap"
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

func (n *NeutralCP) Start(ctx context.Context, logger *zap.Logger) error {
	logger.Info("Start neutral-cp")

	switch n.Config.Registry {
	case PYROSCOPE:
		err := n.startPyroscope(ctx)
		if err != nil {
			return err
		}
	case CLOUD_PROFILER:
		err := n.startCloudProfiler(ctx)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported registry")
	}

	return nil
}
