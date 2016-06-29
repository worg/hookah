![Hoookah](https://dl.dropboxusercontent.com/u/7133562/personal_project/hookah/hookah.min.svg)

# Hookah

A tool written in Go [golang] to handle GitHub and GitLab webhooks [trigger deployments, tests, etc] and
send Telegram notifications


[![GoDoc](https://godoc.org/github.com/worg/hookah/webhooks?status.svg)](https://godoc.org/github.com/worg/hookah)
[![Go Report Card](https://goreportcard.com/badge/github.com/worg/hookah)](https://goreportcard.com/report/github.com/worg/hookah)

## Installing 

Assuming you have a working Go setup

```
go get -u github.com/worg/hookah/cmd/hookah
go install github.com/worg/hookah/cmd/hookah
```

## Usage

You'll need  to create a JSON config file named ```config.json```.

Example:

```
{
    "host": "127.0.0.1",
    "port": 8080,
    "repos": [
        {
            "name": "hookah",
            "branch": "*",
            "tasks" : [
                {
                    "cwd": "/home/user/hookah/",
                    "cmd": "./test.sh",
                    "args": ["prod", "dev"]
                }
            ],
            "notify": {
                "telegram" : {
                    "token": "bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11",
                    "chat_id": 123456
                }
            }
        }
    ]
}
```

### Config Field definition

| field                         | type           | definition                                             | required |
| ---                           | ---            | ---                                                    | ---      |
| host                          | string         | ip/name to listen on                                   | NO       |
| port                          | number         | port to listen on                                      | YES      |
| repos                         | array[object]  | list of repositories to check                          | YES      |
| repos:name                    | string         | repository name                                        | YES      |
| repos:branch                  | string         | repository branch<br>*tip*: you can use * to match any | YES      |
| repos:tasks                   | arrray[object] | list of tasks to execute                               | NO       |
| repos:tasks:cwd               | string         | working directory of command to execute                | NO       |
| repos:tasks:cmd               | string         | command to execute                                     | YES      |
| repos:tasks:args              | string         | arguments passed to command                            | NO       |
| repos:notify                  | object         | notification configuration                             | NO       |
| repos:notify:telegram:token   | string         | Telegram bot token                                     | YES      |
| repos:notify:telegram:chat_id | string         | Telegram Chat ID to send notifications to              | YES      |


For a complete JSON example see [config.sample.json](https://github.com/worg/hookah/blob/master/config.sample.json).

### Running

Execute ```hookah``` on the same path where config.json is or specify its path ```hookah -path="/path/to/config/" ```.

Hookah should start and listen for webhook events on:

* ``/gitlab`` for GitLab Webhooks.
* ``/github`` for GitHub Webhooks.

If any payload matches branch/repo then it will execute the corresponding tasks and [if configured] notify a commit summary.

## Development

The package ```github.com/worg/hookah/webhooks``` contains structures matching GitHub/GitLab payloads for JSON unmarshalling,
both meet Context interface that provides common fields to both payloads.

