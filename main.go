package main

import (
	"flag"
	"github.com/peterbourgon/ff"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fs := flag.NewFlagSet("terrafmter", flag.ExitOnError)
	var (
		resource = fs.String("resource", "nil", "resource to create")
		data     = fs.String("data", "nil", "data source to query")
		file     = fs.String("file", "main.tf", "auto generate terraform file")
		saved    = fs.String("saved", "nil", "saved function")
	)

	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		log.Fatal().Err(err).Msg("parse parameters error")
	}

	if err := Run(*resource, *data, *file, *saved); err != nil {
		log.Fatal().Err(err).Msg("parse terraform templates error")
	}
}

func Run(r, d, filename, saved string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	//s, err := format.Parse(r, d, filename, saved)
	//if err != nil {
	//	return err
	//}
	//if err := gen.Execute(s, os.Stdout); err != nil {
	//	return err
	//}
	return nil
}
