# Scrape Hacker News

1. Create account at [Turso](https://turso.tech):
```bash
turso auth login
turso db create scrape
turso db tokens create scrape
```
Create a URL like:
```
DB_URL=libsql://scrape-jakozaur.turso.io?authToken=[[redacted]]
```
2. Build it:
```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap scraper.go 
zip lambda-handler.zip bootstrap
```
3. Deploy it using hackish ClickOps pratice on AWS:
- https://eu-west-1.console.aws.amazon.com/lambda/home?region=eu-west-1#/functions/scrapeHackerNews?tab=configure
- Use EventBridge to schedule execution every 15 minutes.