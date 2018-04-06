FROM golang:1.10 as build-env

ENV DEP_VERSION v0.4.1
ENV REPOSITORY_PATH /go/src/github.com/matehuszarik/go-api-skeleton

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/${DEP_VERSION}/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p ${REPOSITORY_PATH}
WORKDIR ${REPOSITORY_PATH}

COPY . ${REPOSITORY_PATH}/
RUN dep ensure -vendor-only

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

RUN go build -o server ./main.go
RUN mkdir /app && cp server /app/server

FROM alpine:3.7
COPY --from=build-env /app /app
WORKDIR /app
CMD ["/app/server"]