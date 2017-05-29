# slackPurge

Bulk delete files from slack

## Requirements

- [Go](https://golang.org/dl/)
- [Slack Legacy Token](https://api.slack.com/custom-integrations/legacy-tokens) (For CLI use only)
- Slack team integration token (For slash command integration)

## Building

You can build two ways, one as a CLI tool and one as a cloud deployed team integration.


- To create a CLI to bulk delete files from Slack

Create a `config.toml` in the slackPurge directory like below, replacing the token string with the slack legacy token string
```
# Config file
# Add legacy token here
token = "This is where your token goes"
```

Run `go build`

- Cloud deployed team integration

Create a `token.go` files with your token like the following
```
package main

const token = "This is where your token goes"
```

Run `go build -tags GCE`

## Running

If no errors were thrown you should be able to run the executable with `./slackPurge` 
Runtime options include `-age` in days to delete files that are that many ore more days old and `-c` for the number of files to be deleted at one time.

Run `./slackPurge --help` for more information

For running an integration on AWS or GCP you'll need to upload a zip of the executable and the included `.js` files for GCP Functions(index.js) or AWS Lambda(aws.js)
