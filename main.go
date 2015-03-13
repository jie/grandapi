package main

import (
    "flag"
    "log"
    "./config"
)

func main() {

	port := flag.Int("port", 8080, "server port")
	conf := flag.String("config", "_config/dev.yaml", "Yaml config path")
	debug := flag.Bool("debug", true, "Is debug mode")

	flag.Parse()

	log.Printf(
		"Platapi server will started at: %d, config: %s, debug: %t",
		*port, *conf, *debug)

	_config := Config{}
	c := _config.loadByYaml(*port, *conf, *debug)
	log.Println(c.server.port)
}
