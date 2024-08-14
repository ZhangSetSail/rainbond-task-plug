package config

import (
	"github.com/alecthomas/kingpin/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ESConfig struct {
	EsURL      string
	EsIndex    string
	EsUsername string
	EsPassword string
}

// ParseESFlag -
func ParseESFlag(esc *ESConfig) {
	kingpin.Flag("es-url", "es url").Default(esc.EsURL).Envar("ES_URL").StringVar(&esc.EsURL)
	kingpin.Flag("es-index", "es index").Default(esc.EsIndex).Envar("ES_INDEX").StringVar(&esc.EsIndex)
	kingpin.Flag("es-username", "es username").Default(esc.EsUsername).Envar("ES_USERNAME").StringVar(&esc.EsUsername)
	kingpin.Flag("es-password", "es password").Default(esc.EsPassword).Envar("ES_PASSWORD").StringVar(&esc.EsPassword)
}
