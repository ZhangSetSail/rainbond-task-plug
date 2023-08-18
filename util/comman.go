package util

import (
	"bufio"
	"bytes"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func FileExists(filename string) (bool, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func ExecCommand(command string, args []string, envs []string) error {
	cmd := exec.Command(command, args...)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	cmd.Env = append(os.Environ(), envs...)
	err := cmd.Start()
	if err != nil {
		return err
	}
	logScan := bufio.NewScanner(stdout)
	go func() {
		for logScan.Scan() {
			logrus.Infof(logScan.Text())
		}
	}()

	// 错误日志
	errBuf := bytes.NewBufferString("")
	scan := bufio.NewScanner(stderr)
	for scan.Scan() {
		s := scan.Text()
		logrus.Errorf("exec command filure: %v", s)
		errBuf.WriteString(s)
		errBuf.WriteString("\n")
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
