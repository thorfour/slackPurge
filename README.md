# slackPurge

Bulk delete files from slack

## Requirements

- [Go](https://golang.org/dl/)
- [Slack Legacy Token](https://api.slack.com/custom-integrations/legacy-tokens)

## Building

Create a `config.toml` in the slackPurge directory like below, replacing the token string with the slack legacy token string
```
# Config file
# Add legacy token here
token = "This is where your token goes"
```

Run `go build`

## Running

If no errors were thrown you should be able to run the executable with `./slackPurge` 
Runtime options include `-age` in days to delete files that are that many ore more days old and `-c` for the number of files to be deleted at one time.

Run `./slackPurge --help` for more information
