package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "github.com/rocketlaunchr/google-search"

func main() {
        
	var sel string

	fmt.Println("0: break loop\n1:open webcamxp 5 ip cams")

	looper: for {
		fmt.Scanln(&sel)

		switch sel {
			
		case "0":
			break looper
		case "1":
			fmt.Println(googlesearch.Search(nil, "intitle:webcamxp 5 inurl:8080"))
		case "2":
			fmt.Println(googlesearch.Search(nil, "intitle:webcamxp 6 inurl:8080"))
		}
	}
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}

//search function, i think you can figure it out.
func search(Q int, web, user string) {

        fin := fmt.Sprintf("%s, %s", web, user)
        //search options
        opts := googlesearch.SearchOptions{Limit: Q}
        fmt.Println(googlesearch.Search(nil, fin, opts))

}

func read() string {

        reader := bufio.NewReader(os.Stdin)
        input, er := reader.ReadString('\n')
        check(er)
        input = strings.TrimSuffix(input, "\n")
        //fmt.Println(input)
	return input
}
