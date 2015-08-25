package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func loadConf() {
	var (
		err                  error
		pwd                  string
		configFile, testName string
	)

	if *path != `.` {
		testName = fmt.Sprintf("%s%s%s", *path, PS, cfgName)
		if _, err = os.Stat(testName); err == nil {
			decodeCfg(testName)
		}
	}

	if pwd, err = os.Getwd(); err != nil {
		log.Fatal("ERR: %s", err)
	}

	chunks := strings.Split(pwd, PS)
	cLen := len(chunks)

	for i := range chunks {
		testName = strings.Join(chunks[0:cLen-i], PS) + PS + cfgName
		if _, err = os.Stat(testName); err == nil {
			configFile = testName
			break
		}
	}

	if configFile == `` {
		log.Fatal(`Config file not found… double check path or file names`)
	}

	decodeCfg(configFile)
	logf("Loaded config file file: %s\n", configFile)
}

func decodeCfg(file string) {
	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadFile(file); err != nil {
		log.Fatal(`Config file not found…`)
	}

	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatal(`Error decoding file: `, err)
	}

	logf("%+v", config)
}
