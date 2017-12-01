all:
	go generate .
	gofmt -w *.go
	go install -tags gtk_3_18 
	go build -tags gtk_3_18 
