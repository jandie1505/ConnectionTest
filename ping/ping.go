package ping

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var count = 0

func Application() {

	url := flag.Arg(1)
	logPath := flag.Arg(2)
	interval, err := time.ParseDuration(flag.Arg(3))
	if err != nil {
		log.Fatal(err)
		return
	}

	if interval.Seconds() < 30 {
		log.Fatal("interval must be greater than 30s")
		return
	}

	for {
		ping(url, logPath)
		time.Sleep(interval)
	}
}

func ping(url string, logPath string) {

	status := "OK"
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		status = "FAIL"
	}
	if resp != nil {
		resp.Body.Close()
	}

	timestamp := time.Now().Format("2006-01-02 15-04-05")
	fmt.Println(fmt.Sprintf("%d %s %s %d", count, status, timestamp, url))

	f, _ := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%d %s %s %s\n", count, timestamp, status, url))
	if err != nil {
		log.Fatal(err)
		return
	}

	count++
}
