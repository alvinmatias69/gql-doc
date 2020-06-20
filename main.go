package main

import (
	"os"

	"github.com/alvinmatias69/gql-doc/example"
	"github.com/alvinmatias69/gql-doc/parser"
	"github.com/alvinmatias69/gql-doc/template"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// setup cli args parser
	var (
		cli    = New()
		config = cli.Parse()
	)

	// only show help if help flag is given
	if config.Help || len(os.Args) < 2 {
		cli.ShowHelp()
		return
	}

	// setup logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// only show error on quiet option
	if config.Quiet {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	parserObj, err := parser.New(config.Input, config.Type)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	result, err := parserObj.Parse()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	if !config.NoExample {
		exampleObj := example.New(result, config.Template)
		result, err = exampleObj.Generate()
		if err != nil {
			log.Error().Err(err).Send()
			return
		}
	}

	templating, err := template.New(result, config.Template)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	resultStr, err := templating.Generate()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = write(resultStr, config.Output)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	log.Info().Msg("Documentation generated successfully")
}
