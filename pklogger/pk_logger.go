
// PK custom logger

package pklogger

import (
	"fmt"
	"net"
)

type PKLogger struct {
	conn       *Conn
}

var logger = initLogger()

func initLogger() *PKLogger {
	conn,err := net.Dial("unix", "./pklog.sock")
	if err != nil {
			panic(err.String())
	}

	return &PKLogger{
		conn: conn,
	}
}

func PKLoggerSend(line string) {
	fmt.Println("_PKLOG", line);
}
