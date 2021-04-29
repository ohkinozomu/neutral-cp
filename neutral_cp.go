package neutral_cp

import (
	"context"
	"errors"
	"log"
)

type Config struct {
	Registry        string
	ApplicationName string
	ServerAddress   string
}

type NeutralCP struct {
	Config Config
}

func (n *NeutralCP) Start(ctx context.Context) error {
	log.Println("Start neutral-cp")

	switch n.Config.Registry {
	case "pyroscope":
		n.startPyroscope(ctx)

	default:
		return errors.New("unsupported registry: " + n.Config.Registry)
	}

	return nil
}
