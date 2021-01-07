FROM golang:1.13-alpine as builder

RUN mkdir -p /app/
WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY="http://172.26.1.9:5000"
ENV GOSUMDB="off"

ENV OS=linux
ENV ARCH=amd64
ARG GIT_SHA
ARG VERSION
ARG DATE

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build \
    -a -o bin/playground main.go


FROM alpine:3.12


RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk --no-cache --update add curl ca-certificates

WORKDIR /

COPY --from=builder /app/bin/playground .

CMD ["./playground"]