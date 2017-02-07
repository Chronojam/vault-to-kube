build:
	mkdir -p bin/	
	go build -o bin/kube-test ./kube-test/

release:
	tar -cvzf release-$(shell git rev-parse HEAD).tar.gz bin/*
