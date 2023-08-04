FROM golang:1.19.7-alpine

RUN \
    apk add --no-cache bash git openssh && \
    apk --no-cache add curl && \
    apk --no-cache add vim && \
    apk --no-cache add procps-dev && \
    apk --no-cache add busybox-extras

ADD ./ /app
RUN cd /app
WORKDIR /app
RUN go build -o main ./cmd
CMD ["/app/main"]