ARG GO_VERSION=1.20
ARG ALPINE_VERSION=3.17.2

FROM golang:1.20-alpine AS builder
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
WORKDIR /go/src/project/
COPY . /go/src/project
RUN go mod download
RUN go build -o /bin/project

#FROM alpine:3.17.2
#RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
#WORKDIR /go/src/project/
#COPY --from=build /bin/project /bin/project
#COPY --from=build /go/src/project/.env .

EXPOSE 8080
ENTRYPOINT [ "/bin/project" ]
