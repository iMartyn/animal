FROM golang@sha256:cee6f4b901543e8e3f20da3a4f7caac6ea643fd5a46201c3c2387183a332d989 AS builder
RUN apk update && apk add --no-cache git make ca-certificates && update-ca-certificates
COPY main.go /go/src/github.com/iMartyn/animal/
COPY Makefile /go/src/github.com/iMartyn/animal/
COPY src /go/src/github.com/iMartyn/animal/src
RUN cd /go/src/github.com/iMartyn/animal/; make deps
RUN cd /go/src/github.com/iMartyn/animal/; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o animal .
#RUN ls /go/src/github.com/animal/ -l


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/iMartyn/animal/animal /app/
COPY html /app/html