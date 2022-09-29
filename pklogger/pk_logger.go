
// PK custom logger

package pklogger

import (
	"fmt"
	"net"
	"os"
	"errors"
)

type PKLogger struct {
	listener       net.Listener
	conn           net.Conn
}

var logger *PKLogger

func initLogger() {
	if fileExists("./pklog.sock") {
		os.Remove("./pklog.sock")
	}

	listener,err := net.Listen("unix", "./pklog.sock")
	if err != nil {
			panic(err)
	}

	logger = &PKLogger{
		listener: listener,
	}

	go listenLogger()
}

func fileExists(filePath string) (bool) {
	info, err := os.Stat(filePath)
	if err == nil {
			return !info.IsDir()
	}
	if errors.Is(err, os.ErrNotExist) {
			return false
	}
	return false
}

func listenLogger() {
	for {
		fd, err := logger.listener.Accept()
		if err != nil {
			fmt.Println("PKLOG accept error:", err)
		}

		logger.conn = fd
	}
}

func PKLoggerSend(line string) {
	if logger == nil {
		initLogger()
	}

	//fmt.Println("_PKLOG", line);
	if logger.conn != nil {
		logger.conn.Write([]byte(line))
	}
}
