FROM golang:1.14 AS builder
WORKDIR /build
ADD main.go main.go
ADD go.mod go.mod
ADD cmd cmd
ADD pkg pkg

# Fetch dependencies
RUN go mod download

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /multi-git -a -tags netgo -ldflags '-s -w' .

FROM alpine/git
MAINTAINER Gigi Sayfan <the.gigi@gmail.com>
COPY --from=builder /multi-git /multi-git
ENTRYPOINT ["/multi-git"]
