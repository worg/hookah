package main

const (
	msgTmpl = `
{{.hook.Author.Name}} pushed {{if eq (len .hook.Commits) 0 }}a new branch: {{.branch}}{{else}}{{.hook.Commits | len}} commit[s] to [{{.hook.Repo.Name}}:{{.branch}}]({{.hook.Repo.URL}}/tree/{{.branch}}){{end}}
{{range .hook.Commits}}{{$id := (.ID  | printf "%.7s")}}
    [{{$id}}]({{.URL}}): {{ trimSpace .Message | printf "%.80s"  }}{{if gt (len .Message) 79 }}…{{end}} — {{if .Author.Name}}{{.Author.Name}}{{else}}{{.Author.Username}}{{end}}{{/* 
    no newline between commits
*/}}{{end}}
`
)
