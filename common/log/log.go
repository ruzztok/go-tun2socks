package log

import (
	"net"
	"strconv"

	"github.com/eycorsican/go-tun2socks/common/lsof"
)

var logger Logger

func RegisterLogger(l Logger) {
	logger = l
}

func SetLevel(level LogLevel) {
	if logger != nil {
		logger.SetLevel(level)
	}
}

func Debugf(msg string, args ...interface{}) {
	if logger != nil {
		logger.Debugf(msg, args...)
	}
}

func Infof(msg string, args ...interface{}) {
	if logger != nil {
		logger.Infof(msg, args...)
	}
}

func Warnf(msg string, args ...interface{}) {
	if logger != nil {
		logger.Warnf(msg, args...)
	}
}

func Errorf(msg string, args ...interface{}) {
	if logger != nil {
		logger.Errorf(msg, args...)
	}
}

func Fatalf(msg string, args ...interface{}) {
	if logger != nil {
		logger.Fatalf(msg, args...)
	}
}

func Access(outbound, network, localAddr, target string) {
	localHost, localPortStr, _ := net.SplitHostPort(localAddr)
	localPortInt, _ := strconv.Atoi(localPortStr)
	cmd, err := lsof.GetCommandNameBySocket(network, localHost, uint16(localPortInt))
	if err != nil {
		cmd = "unknown process"
	}
	Infof("[%v] [%v] [%v] %s", outbound, network, cmd, target)
}
