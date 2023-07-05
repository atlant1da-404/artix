package tcp

import (
	"log"
	"net"
)

func NewServer() {

	listen, err := net.Listen("tcp", "localhost"+":"+"5033")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			conn.Close()
			return
		}

		HandleReq(conn)
	}
}
