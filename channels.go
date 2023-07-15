package main

import (
	"fmt"
	"net/http"
	"time"
)

func dispatch() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://staackoverflow.com",
		"http://golnag.org",
		"http://amazon.com",
	}

	ch := make(chan string)

	for _, link := range links {
		// run routine for every call
		go checkLink(link, ch)
	}

	// fmt.Println(<-ch)
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-ch)
	// }
	for l := range ch {
		// time.Sleep(5 * time.Second)
		// go checkLink(l, ch)

		// function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	time.Sleep(5 * time.Second)

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down")
		// ch <- "Might be down"
		ch <- link
		return
	}

	fmt.Println(link, "is up")
	// ch <- "It's up"
	ch <- link
}

// package main

// import "fmt"

// func main() {
//  c := make(chan string)
//  for i := 0; i < 4; i++ {
//      go printString("Hello there!", c)
//  }

//  for s := range c {
//      fmt.Println(s)
//  }
// }

// func printString(s string, c chan string) {
//  fmt.Println(s)
//  c <- "Done printing."
// }

// package main

// import "fmt"

// func main() {
//  c := make(chan string)

//  for i := 0; i < 4; i++ {
//      go printString("Hello there!", c)
//  }

//  for {
//      fmt.Println(<- c)
//  }
// }

// func printString(s string, c chan string) {
//  fmt.Println(s)
//  c <- "Done printing."
// }
