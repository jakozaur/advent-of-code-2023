# Scrape Hacker News with keyword notifications

1. Follow [day two](../02-scrape-in-cloud/README.md).

2. Create account at [Mailgun](https://mailgun.com).

3. Create env file:
```
DB_URL=...
TRIGGERS=your-email@example.com:keyword1,data
MAILGUN_DOMAIN=example.com
MAILGUN_API_KEY=...
MAILGUN_SENDER=hnews@example.com
```

4. You may test it locally, by modifing `lambda.go` to call `HandleRequest`.
```
set -a; source .env; set +a
go run *.go
```

5. Deploy similarly as in [day two](../02-scrape-in-cloud/README.md):
```
GOOS=linux GOARCH=amd64 go build -o bootstrap *.go 
zip lambda-handler.zip bootstrap
```
