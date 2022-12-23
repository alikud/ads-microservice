FROM golang:1.18.1-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY app .
COPY . .
RUN go mod tidy
RUN go build -o /main main.go

#after build
FROM alpine:3
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]