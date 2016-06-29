package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	osPS = string(os.PathSeparator)
)

func loadConf() {
	var (
		err                  error
		pwd                  string
		configFile, testName string
	)

	if *path != `.` {
		testName = fmt.Sprintf("%s%s%s", *path, osPS, cfgName)
		if _, err = os.Stat(testName); err == nil {
			decodeCfg(testName)
			return
		}
	}

	if pwd, err = os.Getwd(); err != nil {
		log.Fatal(`ERR: `, err)
	}

	chunks := strings.Split(pwd, osPS)
	cLen := len(chunks)

	for i := range chunks {
		testName = strings.Join(chunks[0:cLen-i], osPS) + osPS + cfgName
		if _, err = os.Stat(testName); err == nil {
			configFile = testName
			break
		}
	}

	if configFile == `` {
		log.Fatal(`Config file not found… double check path or file names`)
	}

	decodeCfg(configFile)
	log.Printf("Loaded config file file: %s\n", configFile)
}

func decodeCfg(file string) {
	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadFile(file); err != nil {
		log.Fatal(`Config file not found…`)
	}

	if err = json.Unmarshal(data, &cfg); err != nil {
		log.Fatal(`Error decoding file: `, err)
	}
}
