FROM ubuntu:latest

COPY sooty-tern /

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/sooty-tern"]
