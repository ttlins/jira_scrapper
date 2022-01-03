# jira_scrapper

## Description
Simple CLI for scrapping jira data.

All data is cached in json format so jira doesn't need to be queries again for repeated data every time.

## Instructions
- Create folder in your home directory:
```sh
mkdir ~/.jira_scrapper
```

- Copy the example config to the newly created repository and replace the values accordingly
```sh
cp config.yaml.example ~/.jira_scrapper/config.yaml
vi !$
# Change token, host and configure the projects and boards you'd like to query
```

- Compile
```sh
go build
```

## Commands
- boards
    - lists the boards by the projects keys
    - this can be used as a helper for configuration
    - first set the project key, then use this command to get the board_id you'd like to query for
    - creates the cache under `~/.jira_scrapper/cache/boards/`

- sprints
    - must be run before running `issues` (since that will use the generated sprints cache)
    - iterates the projects and retrieves all 'closed' sprints for the configured boards
    - creates the cache under `~/.jira_scrapper/cache/sprints/`

- issues
    - iterates all sprints and returns all issues assigned to each of them
    - creates the cache under `~/.jira_scrapper/cache/issues/`
