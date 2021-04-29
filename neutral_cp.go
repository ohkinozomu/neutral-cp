package neutral_cp

import (
	"context"
	"errors"
	"log"
)

type Config struct {
	registry        string
	applicationName string
	serverAddress   string
}

type NeutralCP struct {
	config Config
}

func (n *NeutralCP) Start(ctx context.Context) error {
	log.Println("Start neutral-cp")

	switch n.config.registry {
	case "pyroscope":
		n.startPyroscope(ctx)

	default:
		return errors.New("unsupported registry: " + n.config.registry)
	}

	return nil
}
