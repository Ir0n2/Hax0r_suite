package main

import (//"io"
	//"net/http"
	"net"
	"fmt"
	//"bufio"
	//"os/exec"
	//"time"
)

func check(e error) {
	if e != nil{
	panic(e)
	}
} 

func main() {

	//listening for connections on port 8080
	listen, err := net.Listen("tcp", ":8080")
	check(err)

	defer listen.Close()
		//conn is the accepted connection
	for {
		conn, er := listen.Accept()
		check(er)

		go handle(conn)

	}
}

func handle(c net.Conn) {
var sel string
	for {
		fmt.Println("type anything to send to the client!")
		//io.Copy(c, c)
		fmt.Scanln(&sel)
		c.Write([]byte(sel))
		/*data, err := bufio.NewReader(c).ReadString('\n')
		check(err)
		fmt.Println(data)*/
	
	}
	c.Close()
}

