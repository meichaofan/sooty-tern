FROM ubuntu:latest

COPY main /

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/main"]
