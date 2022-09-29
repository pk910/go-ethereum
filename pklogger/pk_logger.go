
// PK custom logger

package pklogger

import (
	"fmt"
	"net"
)

type PKLogger struct {
	listener       net.Listener
	conn           net.Conn
}

var logger = initLogger()
go listenLogger()

func initLogger() *PKLogger {
	listener,err := net.Listen("unix", "./pklog.sock")
	if err != nil {
			panic(err)
	}

	return &PKLogger{
		listener: listener,
	}
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
	fmt.Println("_PKLOG", line);
}
