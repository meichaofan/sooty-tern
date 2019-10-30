FROM alpine:latest

COPY sooty-tern /
COPY configs/app.dev.toml /

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/sooty-tern"]
CMD ["-c","app.toml"]
