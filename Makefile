go:
	nodemon --watch './**/*.go' --signal SIGTERM --exec APP_ENV=dev 'go' run main.go

dok:
	docker-compose up --build