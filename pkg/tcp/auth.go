package tcp

import (
	"bufio"
	"net"
	"strings"
)

func login(conn net.Conn) {

	_, _ = conn.Write([]byte("Welcome to Artix Memory Storage! To continue enter username and password! \n"))

	var username string
	var password string

	for strings.TrimSpace(username) != "admin" && strings.TrimSpace(password) != "qwerty" {

		_, _ = conn.Write([]byte("username: "))
		username, _ = bufio.NewReader(conn).ReadString('\n')

		_, _ = conn.Write([]byte("password: "))
		password, _ = bufio.NewReader(conn).ReadString('\n')
	}

	_, _ = conn.Write([]byte("successfully continue to use Artix\n"))
}
