FROM golang:1.19

WORKDIR /gowork

COPY go.mod go.sum ./

ENV GOPROXY=https://proxy.golang.com.cn,direct

RUN go mod download

COPY . ./

RUN go build

EXPOSE 7799

ENTRYPOINT ["./dockerGo"]