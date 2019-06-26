package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Cached struct {
	URL        string
	response   string
	code       int
	lastUpdate int64
	headers    http.Header
}

type Endpoint struct {
	URL      string `json:"url"`
	Method   string `json:"method"`
	Interval int64  `json:"interval"`
	cached   []*Cached
}

type Config struct {
	PublicPort   string `json:"PublicPort"`
	PrivatePort  string
	PrivateURL   string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Enpoints     []*Endpoint `json:"enpoints"`
}

func loadConfig() {
	loadFile()
	if config.PublicPort == "" {
		config.PublicPort = "8000"
	}
	if config.PrivatePort == "" {
		config.PrivatePort = "80"
	}
	if config.PrivateURL == "" {
		config.PrivateURL = "localhost"
	}
	if config.WriteTimeout == 0 {
		config.WriteTimeout = 30
	}
	if config.ReadTimeout == 0 {
		config.ReadTimeout = 30
	}
}

func loadFile() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	config = &Config{}
	config.Enpoints = make([]*Endpoint, 0)
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Panic(err.Error())
	}
}
