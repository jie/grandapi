package config

import "io/ioutil"
import "path/filepath"
import "gopkg.in/yaml.v2"


type Config struct {
	server struct {
		port  int
		debug bool
	}

    database struct {
        name string
        conn string
    }

    redis struct {
        host string
        port int
        db int
        timeout int
    }
}

func (config Config) loadByYaml(port int, configfile string, debug bool) Config {
	filename, _ := filepath.Abs(configfile)
	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	config.server.port = port
	config.server.debug = debug
	return config
}
