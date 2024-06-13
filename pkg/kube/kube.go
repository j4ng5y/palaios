package kube

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, args []string) error {
	llvl, err := cmd.PersistentFlags().GetString("log-level")
	if err != nil {
		return fmt.Errorf("unable to retrieve log-level: %w", err)
	}
	lvl, err := zerolog.ParseLevel(llvl)
	if err != nil {
		return fmt.Errorf("unable to parse provided log level: \"%s\": %w", llvl, err)
	}

	zerolog.SetGlobalLevel(lvl)

	logger := zerolog.Ctx(cmd.Context())

	logger.Debug().Msg("successfully initialized logger")
	return nil
}
