FROM golang:1.11.13-alpine3.10
LABEL maintainer="Eric Hartline <eric.hartline@northwestern.edu>"
RUN apk add --no-cache git build-base
RUN mkdir -p /file-uploader
WORKDIR /file-uploader
ADD cmd cmd
ADD internal internal
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o app cmd/file-uploader/main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=0 /file-uploader/app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]
