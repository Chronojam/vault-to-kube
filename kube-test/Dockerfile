FROM debian:7
MAINTAINER calum

RUN wget https://github.com/Chronojam/vault-to-kube/releases/download/v0.0.1/vault-to-kube-v0.0.1.tar.gz
RUN tar -xzf vault-to-kube-v0.0.1.tar.gz

ENTRYPOINT ["/kube-test"]
