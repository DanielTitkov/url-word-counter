package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func countTokenEntries(str, token string) (int, error) {
	re, err := regexp.Compile(token)
	if err != nil {
		return 0, err
	}
	matches := re.FindAllString(str, -1)
	return len(matches), nil
}

func countTokenAtURL(url, token string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	n, err := countTokenEntries(string(body), token)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func main() {
	token := flag.String("token", "Go", "Token to count (regexp supported)")
	jobs := flag.Int("jobs", 5, "Max concurrent jobs")
	flag.Parse()

	log.Printf("Starting with max %d jobs", *jobs)

	sem := make(chan struct{}, *jobs)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		sem <- struct{}{}
		go func(url, token string, sem chan struct{}) {
			defer func() { <-sem }()
			log.Printf("Started processing url '%s'", url)
			n, err := countTokenAtURL(url, token)
			if err != nil {
				log.Printf("Failed to process url '%s': %v", url, err)
			}
			log.Printf("Got result for '%s': %d entries of token '%s'", url, n, token)
		}(url, *token, sem)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error occured: %v", err)
	}

	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
}
