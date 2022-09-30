
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
	listener,err := net.Listen("tcp", "0.0.0.0:8605")
	if err != nil {
			return
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
		_, err := logger.conn.Write([]byte(line))
		if err != nil {
      logger.conn.Close() // close if problem
      logger.conn = nil
   	}
	}
}
