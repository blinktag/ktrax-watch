package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/caarlos0/env"
)

type config struct {
	IDs   []string `env:"IDS" envSeparator:":"`
	Found []string
}

func main() {

	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	ticker := time.NewTicker(time.Millisecond * 60000)

	for _ = range ticker.C {
		search(cfg)
	}
}

func search(cfg config) {
	url := "http://ktrax.kisstech.ch/cgi-bin/flarm-txrange.cgi?command=who"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	body := string(html)

	for i, id := range cfg.IDs {
		if strings.Contains(body, id) {
			cfg.Found = append(cfg.Found, id)
			cfg.IDs = append(cfg.IDs[:i], cfg.IDs[i+1:]...)
			fmt.Printf("Found %s!\n", id)
		}
	}
}
