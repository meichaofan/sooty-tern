FROM alpine:latest

COPY sooty-tern /

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/sooty-tern"]
CMD ["-c","configs/app.toml"]
