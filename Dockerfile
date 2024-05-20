FROM golang:1.22.3 AS builder
LABEL maintainer="hugh.12345zxcvb@gmail.com"
ENV GIN_MODE=release
WORKDIR /src/
COPY . /src
RUN set -eux;\
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app ./cmd

FROM scratch
COPY --from=builder /src/app /bin/app
CMD ["/bin/app"]