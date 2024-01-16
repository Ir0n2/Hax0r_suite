package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

//var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
//var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

func main() {

	var sel string 
	var wordlist string
	var md5hash string

	var sha256hash string
	fmt.Println("1: MD5\n2:SHA256")
	fmt.Scanln(&sel)

	if sel == "1" {
		fmt.Println("md5 hash:")
		fmt.Scanln(&md5hash)
	} else {
		fmt.Println("sha256 hash:")
		fmt.Scanln(&sha256hash)
	}

	fmt.Println("Please paste wordlist url")
	fmt.Scanln(&wordlist)


	f, err := os.Open(wordlist)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()

		switch sel {
			case "1":
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == md5hash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
			}
			case "2":
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sha256hash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
