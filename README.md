# Gator

## Requirements
- Postgres
- Go

## Installation
Install the CLI with:

```bash
go install github.com/ssgkian/gator@latest
```

## Config

Create a `.gatorconfig.json` file in your home directory.

```json
{
    "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

## Usage

gator register <name>
gator login <name>
gator reset
gator addfeed <name> <url>
gator feeds
gator follow <url>
gator unfollow <url>
gator agg <time_duration> (30s, 1m, 15m etc)
gator browse <limit> (5, 10, 20 etc)
