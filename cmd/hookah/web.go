package main

import (
	"bytes"
	"encoding/json"
	"github.com/tucnak/telebot"
	"github.com/worg/hookah/webhooks"
	"log"
	"net/http"
	"strings"
	"text/template"
)

const (
	msgTmpl = `
{{.hook.Author.Name}} pushed {{if eq (len .hook.Commits) 0 }}a new branch: {{.branch}}{{else}}{{.hook.Commits | len}} commit[s] to {{.hook.Repo.Name}}:{{.branch}}{{end}}
{{range .hook.Commits}}
    {{.ID |printf "%.7s"}}: {{ trimSpace .Message | printf "%.80s"  }}{{if gt (len .Message) 79 }}…{{end}} — {{if .Author.Name}}{{.Author.Name}}{{else}}{{.Author.Username}}{{end}}{{/* 
    no newline between commits
*/}}{{end}}`
)

var (
	tmpl *template.Template
)

func init() {
	tmpl = template.Must(template.New(`pushMsg`).Funcs(template.FuncMap{
		`trimSpace`: strings.TrimSpace,
	}).Parse(msgTmpl))
}

func gitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != `POST` {
		http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	switch strings.TrimPrefix(r.URL.String(), `/`) {
	case `gitlab`:
		var hook webhooks.GitLab

		if err := decoder.Decode(&hook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		processHook(hook)
	case `github`:
		var hook webhooks.GitHub

		switch r.Header.Get(`X-GitHub-Event`) {
		case `push`:
			break
		case `ping`: // just return on ping
			w.WriteHeader(http.StatusOK)
			return
		default:
			http.Error(w, ``, http.StatusNotAcceptable)
			return
		}

		if err := decoder.Decode(&hook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		processHook(hook)
	}

	w.WriteHeader(http.StatusOK)
}

func processHook(ctx webhooks.Context) {
	h := ctx.Hook()
	branch := strings.TrimPrefix(h.Ref, `refs/heads/`)
	for _, r := range cfg.Repos {
		go func(r repo) {
			if r.Name != h.Repo.Name ||
				(r.Branch != `*` && r.Branch != branch) {
				return
			}

			go r.Tasks.Run() //execute tasks
			if r.Notify.Telegram.ChatID != 0 &&
				r.Notify.Telegram.Token != `` {
				var (
					buf bytes.Buffer
					bot *telebot.Bot
					err error
				)

				err = tmpl.Execute(&buf, map[string]interface{}{
					`hook`:   h,
					`branch`: branch,
				})
				if err != nil {
					log.Println(`Template ERR:`, err)
					return
				}

				if bot, err = telebot.NewBot(r.Notify.Telegram.Token); err != nil {
					log.Println(`Telegram ERR:`, err)
					return
				}

				err = bot.SendMessage(telebot.User{ID: r.Notify.Telegram.ChatID}, buf.String(), nil)
				if err != nil {
					log.Println(`Telegram ERR:`, err)
					return
				}

				log.Println(`Message Sent`)
			}
		}(r)
	}
}
