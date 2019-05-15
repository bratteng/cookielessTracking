FROM golang:alpine AS builder
WORKDIR /go/src/cookielessTracking
COPY main.go .

ENV PORT=9000

RUN set -xe; \
	\
	apk update; \
	apk add --no-cache --virtual .build-deps \
		git \
	; \
	\
	go get -u github.com/julienschmidt/httprouter; \
	go get -u github.com/caarlos0/env; \
	go get -u github.com/satori/go.uuid; \
	\
	apk del .build-deps; \
	rm -rf /var/cache/apk/*

RUN set -xe; \
	\
	go build; \
	ls -la

FROM alpine
COPY --from=builder /go/src/cookielessTracking/cookielessTracking /
CMD ["/cookielessTracking"]
