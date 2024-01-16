//GoSniffer is a social media account sniffing application

package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "github.com/rocketlaunchr/google-search"

func main() {
	var fuk int
	fmt.Println("Type a username and it will make multiple searches on various websites")
        input := read()
	fmt.Println("type how many results it will show, If there aren't enough, it will only print the ones avaliable.")
	fmt.Scanln(&fuk)
        fmt.Println("Youtube")
        search(fuk, "site:www.youtube.com", input)
        fmt.Println("Instagram")
        search(fuk, "site:www.instagram.com", input)
        fmt.Println("Twitch")
        search(fuk, "site:www.twitch.com", input)
        fmt.Println("Xbox")
        search(fuk, "site:www.xbox.com", input)
        fmt.Println("Reddit")
        search(fuk, "www.reddit.com", input)
	fmt.Println("tiktok")
        search(fuk, "www.tiktok.com", input)
	fmt.Println("Tumblr")
        search(fuk, "www.tumblr.com", input)
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


