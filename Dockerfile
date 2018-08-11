FROM golang:alpine as builder

RUN apk add --no-cache make \
	git

RUN go get -u github.com/golang/dep/cmd/dep && \
	go get -u github.com/beito123/medaka

WORKDIR /go/src/github.com/beito123/medaka

RUN make deps && \
	make

FROM alpine:latest
COPY --from=builder /go/src/github.com/beito123/medaka/medaka ./
CMD ./medaka