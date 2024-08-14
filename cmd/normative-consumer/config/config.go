package config

import "C"
import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/goodrain/rainbond-task-plug/pkg/config"
)

var s *NormativeServer

func init() {
	s = &NormativeServer{
		Config: Config{
			NatsHost: "127.0.0.1",
			NatsPort: "4222",
			Port:     "10010",
			DB: &config.DBConfig{
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

// NormativeServer  normative server
type NormativeServer struct {
	Config
	LogLevel string
}

type Config struct {
	NatsHost string
	NatsPort string
	Port     string
	DB       *config.DBConfig
}

func Parse() {
	kingpin.Flag("log-level", "The level of logger").Default(s.LogLevel).Envar("LOG_LEVEL").StringVar(&s.LogLevel)
	kingpin.Flag("nats-host", "nats host:127.0.0.1").Default(s.NatsHost).Envar("NATS_HOST").StringVar(&s.NatsHost)
	kingpin.Flag("nats-port", "nats port:4222").Default(s.NatsPort).Envar("NATS_PORT").StringVar(&s.NatsPort)
	kingpin.Flag("port", "server port:10010").Default(s.NatsPort).Envar("PORT").StringVar(&s.Port)
	config.ParseDBFlag(s.DB)
	kingpin.CommandLine.GetFlag("help").Short('h')
	kingpin.Parse()
}

func GetNormativeConsumerServer() *NormativeServer {
	return s
}
