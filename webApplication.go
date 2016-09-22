package main

import (

    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    "bufio"
    "strings"
)

func main() {

	channel := make(chan string)

	arr := make([]string, 0)
	fmt.Print("Enter text: ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
            break
        }
		arr = append(arr, line)
	}
	for _,a := range arr{
		  go getHttpResponse(a, channel) //goroutine
	}   

    for range arr{
        fmt.Println(<-channel) // receive 
    }

}

func getHttpResponse(url string, channel chan<- string){
		
		url = strings.TrimRight(url, "\r\n")
		url = strings.TrimSpace(url)
	
		response, err := http.Get(url)

        if err != nil {
            channel <- fmt.Sprintf("error while requesting: %v\n", err)
            return
        }

        bodyContent, err := ioutil.ReadAll(response.Body)

        response.Body.Close()

        if err != nil {
            channel <- fmt.Sprintf( "error while reading  %s: %v\n", url, err)
            return
        }

        channel <- fmt.Sprintf("%s\n : %s\n", url,bodyContent)

}