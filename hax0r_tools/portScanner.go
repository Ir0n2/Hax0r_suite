package main

import ("fmt"
	"net"
)

func main() {
/*	var fuck string
	
	var port int
	
	fmt.Println("Wesbite or Ip address")

	fmt.Scanln(&fuck)
	
	fmt.Println("Scan up to port:")

        fmt.Scanf(&port)*/


	for i := 10; i <= 90; i++ {
			
		address := fmt.Sprintf("nmap.org:%d", i)

		conn, err := net.Dial("tcp", address)

		if err == nil {
			fmt.Println("connection")
			conn.Close()
		}

		if err != nil {
			fmt.Println("closed")
			continue
		}
	}
}
