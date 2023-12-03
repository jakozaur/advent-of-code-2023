# Scrape Hacker News

Uses [Colly](https://go-colly.org/).

To run:
```bash
  go run scraper.go
  sqlite3 hackernews.db
  sqlite> .schema submissions 
```

