package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const BaseURL = "https://api.openweathermap.org"

func main() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
		os.Exit(1)
	}
	URL := fmt.Sprintf("%s/data/2.5/weather?q=London,UK&appid=%s", BaseURL, key)
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "unexpected response status", resp.Status)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
}
