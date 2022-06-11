FROM golang:1.16-alpine

WORKDIR /app
ARG GOPROXY

ENV GOPROXY="https://proxy.golang.com.cn,direct"
ENV GOARCH=$GOARCH
ENV GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=1

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-web-server

EXPOSE 8080

CMD [ "/docker-web-server" ]
