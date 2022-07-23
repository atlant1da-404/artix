package tcp

import (
	"bufio"
	"encoding/json"
	"github.com/atlant1da-404/artik_db/pkg/bucket"
	"net"
	"strings"
)

func requests(conn net.Conn) {

	buck := bucket.New()

	for {
		str, _ := bufio.NewReader(conn).ReadString('\n')

		word := strings.Fields(str)
		var response interface{}

		switch strings.ToLower(word[0]) {
		case "get":

			if err := buck.Get(word[1], &response); err != nil {
				_, _ = conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			value, _ := json.Marshal(response)
			_, _ = conn.Write(value)
			_, _ = conn.Write([]byte("\n"))

		case "insert":

			if err := buck.Insert(word[1], []byte(word[2])); err != nil {
				_, _ = conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			_, _ = conn.Write([]byte("Inserted!\n"))

		case "dump":

			if err := bucket.Dump("dump.json", *buck); err != nil {
				_, _ = conn.Write([]byte(err.Error() + "\n"))
				continue
			}

		case "load_dump":

			if err := bucket.LoadDump("dump.json", buck); err != nil {
				_, _ = conn.Write([]byte(err.Error() + "\n"))
				continue
			}
		case "update":

			if err := buck.Update(word[1], []byte(word[2])); err != nil {
				_, _ = conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			_, _ = conn.Write([]byte("1" + "\n"))

		case "del":

			if err := buck.Delete(word[1]); err != nil {
				_, _ = conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			_, _ = conn.Write([]byte("1" + "\n"))

		case "exit":
			return

		default:
			_, _ = conn.Write([]byte("unknown command\n"))
		}
	}
}
