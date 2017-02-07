build:
	mkdir -p bin/	
	go build -o bin/kube-test ./kube-test/

release:
	tar -cvzf vault-to-kube-$(shell cat VERSION).tar.gz bin/*
