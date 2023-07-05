package tcp

import (
	"net"
)

func HandleReq(conn net.Conn) {
	defer conn.Close()
	login(conn)
	requests(conn)
}
