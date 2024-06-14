package config

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/goodrain/rainbond-task-plug/pkg"
)

var p *ProducerServer

func init() {
	p = &ProducerServer{
		Config: Config{
			NatsHost: "127.0.0.1",
			NatsPort: "4222",
			Port:     "10010",
			DB: &pkg.DBConfig{
				DBName:       "taskPlug",
				DBUser:       "root",
				DBPass:       "password",
				DBHost:       "127.0.0.1",
				DBPort:       "3306",
				MaxOpenConns: 1024,
			},
		},
		LogLevel: "info",
	}
}

// ProducerServer  producer server
type ProducerServer struct {
	Config
	LogLevel string
}

type Config struct {
	NatsHost string
	NatsPort string
	Port     string
	DB       *pkg.DBConfig
}

func Parse() {
	kingpin.Flag("log-level", "The level of logger").Default(p.LogLevel).Envar("LOG_LEVEL").StringVar(&p.LogLevel)
	kingpin.Flag("nats-host", "nats host:127.0.0.1").Default(p.NatsHost).Envar("NATS_HOST").StringVar(&p.NatsHost)
	kingpin.Flag("nats-port", "nats port:4222").Default(p.NatsPort).Envar("NATS_PORT").StringVar(&p.NatsPort)
	kingpin.Flag("port", "server port:10010").Default(p.NatsPort).Envar("PORT").StringVar(&p.Port)

	pkg.ParseDBFlag(p.DB)

	kingpin.CommandLine.GetFlag("help").Short('h')
	kingpin.Parse()
}

func GetProducerServer() *ProducerServer {
	return p
}
