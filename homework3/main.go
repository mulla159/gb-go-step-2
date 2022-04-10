package main

import (
	"flag"

	conf "github.com/mulla159/yaml-conf"
)

func main() {
	config := conf.AppConfig{}

	var fileName = flag.String("fn", "conf.yaml", "File name (path)")
	flag.Parse()

	config.Get(*fileName)

	config.Print()
}
