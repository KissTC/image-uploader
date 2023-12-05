build:
	@go build -o myapp

run: build
	@./myapp