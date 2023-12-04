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
2. 