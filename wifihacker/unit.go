package main

import ("fmt"
	"os"
	"bufio"
	//"strings"
)

func main() {

	fmt.Println(readList("/home/zerocool/wifihacker/test.txt", 1))
	
	

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
