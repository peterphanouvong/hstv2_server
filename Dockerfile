FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/peterphanouvong/hst/
WORKDIR /go/src/github.com/peterphanouvong/hst
RUN go mod download
COPY . /go/src/github.com/peterphanouvong/hst
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/hst github.com/peterphanouvong/hst

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/peterphanouvong/hst/build/hst /usr/bin/hst
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/hst"]