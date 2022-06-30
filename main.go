package main

import (
	"flag"
	"os"

	"github.com/Pangjiping/terrafmtter/util"
	"github.com/peterbourgon/ff"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// todo: more short?
	fs := flag.NewFlagSet("terrafmtter", flag.ExitOnError)
	var (
		resource = fs.String("resource", "nil", "resource to create")
		data     = fs.String("data", "nil", "data source to query")
		version  = fs.String("version", "latest", "terrform provider version for alicloud")
		file     = fs.String("file", "main.tf", "auto generate terraform file")
	)

	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		log.Fatal().Err(err).Msg("parse parameters error")
	}

	resources := util.ConvertString2slice(*resource, ",")
	datas := util.ConvertString2slice(*data, ",")

	if err := Run(resources, datas, *file, *version); err != nil {
		log.Fatal().Err(err).Msg("parse terraform templates error")
	}
}

func Run(resources []string, datas []string, filename, version string) error {
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
