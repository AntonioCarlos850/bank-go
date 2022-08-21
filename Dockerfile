FROM golang

ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    go get github.com/labstack/echo/v4 && \
    go get github.com/mattn/go-sqlite3

CMD ["tail", "-f", "/dev/null"]