FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/go-deepread

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./conf /app/conf
RUN go build -ldflags="-s -w" -o /app/go-deepread main.go


FROM alpine

RUN echo http://mirrors.aliyun.com/alpine/v3.10/main/ > /etc/apk/repositories && \
    echo http://mirrors.aliyun.com/alpine/v3.10/community/ >> /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
RUN mkdir log
COPY --from=builder /app/go-deepread /app/go-deepread
COPY --from=builder /app/conf /app/conf

CMD ["./go-deepread"]
