package main

import (
	"flag"
	"os"

	"github.com/Pangjiping/terrafmtter/gen"

	"github.com/Pangjiping/terrafmtter/util"
	"github.com/peterbourgon/ff"
	"github.com/rs/zerolog/log"
)

func main() {

	// todo: more short?
	fs := flag.NewFlagSet("terrafmtter", flag.ExitOnError)
	var (
		resource = fs.String("resource", "nil", "resource to create")
		data     = fs.String("data", "nil", "data source to query")
		version  = fs.String("version", "latest", "terrform provider version for alicloud")
		file     = fs.String("file", "main.tf", "auto generate terraform file")
	)

	util.Logger.Info().Msg("[INFO] terrafmtter parse parameters...")
	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("[ERROR] parse parameters error")
	}

	resources := util.ConvertString2slice(*resource, ",")
	datas := util.ConvertString2slice(*data, ",")

	util.Logger.Info().Msg("[INFO] terrafmtter run ...")
	if err = Run(resources, datas, *file, *version); err != nil {
		log.Fatal().Err(err).Msg("[ERROR] parse terraform templates error")
	}
	util.Logger.Info().Msg("[INFO] terrafmtter exit. please check target file.")
}

func Run(resources []string, datas []string, filename, version string) error {
	if err := gen.Execute(resources, datas, filename, version); err != nil {
		return err
	}
	return nil
}
