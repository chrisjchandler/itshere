package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"
)

const (
    apiURL = "https://api.com/api/function"
    dir = "/path/to/directory"
)

func main() {
    for {
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            log.Fatal(err)
        }

        for _, file := range files {
            if !file.IsDir() {
                // Call the API with the file path
                resp, err := http.Post(apiURL, "application/json", strings.NewReader(file.Name()))
                if err != nil {
                    log.Printf("Error calling API: %v", err)
                    continue
                }
                defer resp.Body.Close()

                // Handle the response from the API
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    log.Printf("Error reading response: %v", err)
                    continue
                }
                log.Println(string(body))
            }
        }

        time.Sleep(1 * time.Second)
    }
}
//time sleep can be adjusted to your monitoring interval
