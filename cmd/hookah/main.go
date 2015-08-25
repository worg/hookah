package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type (
	Config struct {
		// host to listen on
		Host string `json:"host"`
		// port to listen on
		Port int `json:"port"`
		// list of repos to check
		Repos []repo `json:"repos"`
	}

	repo struct {
		Name   string   `json:"name"`
		Branch string   `json:"branch"`
		Tasks  taskList `json:"tasks"`
		Notify notify   `json:"notify"`
	}

	// notify may contain alternative notifications
	notify struct {
		Telegram telegramBot `json:"telegram"`
	}

	// task is used to specify the commands to execute
	task struct {
		Args []string `json:"args"`
		Cmd  string   `json:"cmd"`
		Cwd  string   `json:"cwd"`
	}

	// telegramBot holds info to connect to telegram bot API
	telegramBot struct {
		ChatID int    `json:"chat_id"`
		Token  string `json:"token"`
	}

	taskList []task
)

const (
	cfgName     = `config.json`
	defaultPort = 8080
	PS          = string(os.PathSeparator)
)

var (
	path    = flag.String(`path`, `.`, `where to look for config file`)
	config  Config
	mux     = http.DefaultServeMux
	sprintf = fmt.Sprintf //share sprintf across files
)

func main() {
	flag.Parse()
	loadConf()

	// Add handlers
	mux.HandleFunc(`/gitlab`, gitHandler)
	mux.HandleFunc(`/github`, gitHandler)

	address := sprintf("%s:%d", config.Host, config.Port)
	log.Printf("Listening on %s/gitlab and %[1]s/github", address)
	log.Fatal("ERROR SERVING: %+v", http.ListenAndServe(address, nil))
}

// Run method for task
func (t task) Run() {
	var (
		out []byte

		err error
	)

	cmd := exec.Command(t.Cmd, t.Args...)
	if cmd.Dir = t.Cwd; cmd.Dir == `` {
		cmd.Dir, _ = os.Getwd() //fallback to current working dir
	}

	if out, err = cmd.CombinedOutput(); err != nil {
		logf("ERROR: %s -- RUNNING: %s %s WITH ARGS: %+v", err, t.Cwd, t.Cmd, t.Args)
	}
	logf("OUTPUT: %v", out)
}

func (tl taskList) Run() {
	for _, t := range tl {
		t.Run()
	}
}

func logf(f string, args ...interface{}) {
	log.Printf("%s%s%[1]s", "\n-----\n", sprintf(f, args...))
}
