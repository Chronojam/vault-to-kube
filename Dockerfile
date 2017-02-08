FROM ubuntu
MAINTAINER calum

RUN apt-get update
RUN apt-get install -y ca-certificates

COPY vault-to-kube /vault-to-kube
COPY cert.pem /usr/local/share/ca-certificates/cert.crt

RUN update-ca-certificates
#RUN cat /etc/vault/cert.pem >> /etc/ssl/certs/ca-certificates.pem

ENTRYPOINT ["/vault-to-kube"]
