package cmd

import (
	"context"
	"os"

	"github.com/j4ng5y/palaios/pkg/kube"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	root = &cobra.Command{
		Use:   "palaios",
		Short: "Palaios is a Kubernetes API deprecation inventory tool.",
		RunE:  kube.Cmd,
	}
)

func Execute() {
	root.PersistentFlags().StringP("kube-config", "c", "", "The Kuberentes config file to use.")
	root.PersistentFlags().StringP("kube-context", "x", "", "The Kubernetes context to use.")
	root.PersistentFlags().StringP("log-level", "l", "info", "The log level to use. [trace|debug|info|warn|error|fatal|panic]")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = zerolog.New(os.Stdout).With().Caller().Logger().WithContext(ctx)

	if err := root.ExecuteContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("unable to execute root command")
	}
}
