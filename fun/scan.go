package fun

import (
	"fmt"
	"net"
)

func funReader() {
	IP := "scanme.nmap.org"

	// 1-10
	for i := 1; i < 10; i++ {
		address := fmt.Sprintf(IP + ":%d", i)

		conn, err := net.Dial("tcp", address)

		if err != nil {
			// handle error
			fmt.Println(err)
		}

		fmt.Printf("[+] Connection established.. PORT: %v\n Address: %s", i, conn.RemoteAddr().String())

		err = conn.Close()

		if err != nil {
			fmt.Println(err)
		}
	}
}