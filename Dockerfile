FROM alpine:latest

COPY sooty-tern /

ENV sooty_tern_env prod

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/sooty-tern"]
CMD ["-c","app.toml"]
