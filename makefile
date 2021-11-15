# go build command
gobuild:
	@go build -v -o main *.go 

gorun:
	make gobuild
	@./main 