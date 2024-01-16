package main

import ("fmt"
        "net"
)

func main() {
        var fuck string
        
        var port int
        
        fmt.Println("Wesbite or Ip address")

        fmt.Scanln(&fuck)
        
        fmt.Println("Port number:")

        fmt.Scanln(&port)

	i := 0

        for {

		i++

		fmt.Println(i)

                address := fmt.Sprintf("%s:%d", fuck, port)

                conn, err := net.Dial("udp", address)

                if err == nil {
                        fmt.Println("Sent")
                        //conn.Close()
                }

                if err != nil {
                        fmt.Println("closed")
                        conn.Close()
			break
                }
        }
}

