FROM golang:1.16 AS builder

ENV WORKDIR=/go/src/app

COPY . $WORKDIR

RUN cd $WORKDIR && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -a -installsuffix cgo -o app ./example.go


FROM scratch

COPY --from=builder /go/src/app/app /

ENTRYPOINT ["/app"]
