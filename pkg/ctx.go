package pkg

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func InitCTX(s time.Duration) {
	logrus.Infof("init ctx")
	ctx, cancel = context.WithTimeout(context.Background(), s*time.Second)
}

func GetCTX() context.Context {
	return ctx
}

func CloseCTX() {
	cancel()
}
