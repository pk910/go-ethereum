
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

var logger PKLogger

func initLogger() {
	listener,err := net.Listen("unix", "./pklog.sock")
	if err != nil {
			panic(err)
	}

	logger = &PKLogger{
		listener: listener,
	}

	go listenLogger()
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

	fmt.Println("_PKLOG", line);
	if logger.conn != nil {
		logger.conn.Write([]byte(line))
	}
}
