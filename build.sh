#!/usr/bin/env bash
go build -o ./app main.go ./database/database.go ./docs/docs.go ./email/email.go ./google_crawler/google_crawler.go ./handlers/identify_key_words.go ./models/models.go
