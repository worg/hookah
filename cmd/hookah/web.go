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

var (
	tmpl *template.Template

	// telegram bot send options
	sendOpts = telebot.SendOptions{
		ParseMode: `Markdown`,
	}

	// functions available on templates
	tFuncs = template.FuncMap{
		`trimSpace`: strings.TrimSpace,
		// backtick message to avoid markdown parsing errors
		`fmtCommit`: func(s string) string {
			return "`" + s + "`"
		},
	}
)

func init() {
	tmpl = template.Must(template.New(`pushMsg`).Funcs(tFuncs).Parse(msgTmpl))
}

func gitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != `POST` {
		http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	switch r.URL.Path {
	case `/gitlab`:
		var hook webhooks.GitLab

		if err := decoder.Decode(&hook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		processHook(hook)
	case `/github`:
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

				err = bot.SendMessage(telebot.User{ID: r.Notify.Telegram.ChatID}, buf.String(), &sendOpts)
				if err != nil {
					log.Println(`Telegram ERR:`, err)
					return
				}

				log.Println(`Message Sent`)
			}
		}(r)
	}
}
