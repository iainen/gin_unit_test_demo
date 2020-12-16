FROM golang:1.13 as builder

WORKDIR /workspace

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o manager cmd/demo/main.go

# FROM gcr.io/distroless/static:nonroot
FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

ENV TZ Asia/Shanghai

WORKDIR /

COPY --from=builder /workspace/manager .

ENTRYPOINT ["/manager"]
