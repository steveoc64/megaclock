all:
	rm -f binary/*
	gofmt -w *.go
	go install -tags gtk_3_18 
	go build -tags gtk_3_18 
	cp megaclock binary
	strip binary/megaclock
	mv binary/megaclock binary/megaclock-linux-x86_64
