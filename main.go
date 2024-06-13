package main

import (
	"github.com/j4ng5y/palaios/cmd"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	cmd.Execute()
}
