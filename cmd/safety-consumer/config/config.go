package config

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/goodrain/rainbond-task-plug/pkg"
)

var c *SafetyConsumerServer

func init() {
	c = &SafetyConsumerServer{
		Config: Config{
			NatsHost:        "127.0.0.1",
			NatsPort:        "4222",
			CodeStoragePath: "/usr/src/",
			SonarToken:      "",
			SonarHost:       "127.0.0.1",
			SonarPort:       "9000",
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

// Config config
type Config struct {
	NatsHost        string
	NatsPort        string
	CodeStoragePath string
	SonarToken      string
	SonarHost       string
	SonarPort       string
	DB              *pkg.DBConfig
}

// SafetyConsumerServer  consumer server
type SafetyConsumerServer struct {
	Config
	LogLevel string
}

func Parse() {
	kingpin.Flag("log-level", "The level of logger").Default(c.LogLevel).Envar("LOG_LEVEL").StringVar(&c.LogLevel)
	kingpin.Flag("nats-host", "nats host:127.0.0.1").Default(c.NatsHost).Envar("NATS_HOST").StringVar(&c.NatsHost)
	kingpin.Flag("nats-port", "nats port:4222").Default(c.NatsPort).Envar("NATS_PORT").StringVar(&c.NatsPort)
	kingpin.Flag("code-storage-path", "code storage address").Default(c.CodeStoragePath).Envar("CODE_STORAGE_PATH").StringVar(&c.CodeStoragePath)
	kingpin.Flag("sonar-token", "sonar token").Default(c.SonarToken).Envar("SONAR_TOKEN").StringVar(&c.SonarToken)
	kingpin.Flag("sonar-host", "sonar host").Default(c.SonarHost).Envar("SONAR_HOST").StringVar(&c.SonarHost)
	kingpin.Flag("sonar-port", "sonar port").Default(c.SonarPort).Envar("SONAR_PORT").StringVar(&c.SonarPort)
	pkg.ParseDBFlag(c.DB)
	kingpin.CommandLine.GetFlag("help").Short('h')

	kingpin.Parse()
}

// GetSafetyConsumerServer get consumer server
func GetSafetyConsumerServer() *SafetyConsumerServer {
	return c
}
