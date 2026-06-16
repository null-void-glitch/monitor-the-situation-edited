package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/radar/annotations/outages?dateRange=1d", nil)
	req.Header.Set("Authorization", "Bearer cfat_RrGgX1Tv0cDEcFMcEhjMhRCGBP3mOQ3YhRT6uzNmb0295523")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	
	fmt.Println(string(b))
}
