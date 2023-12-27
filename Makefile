run:
	go run cmd/*.go

cp-env:
	cp .env.default .env

.PHONY: run cp-env