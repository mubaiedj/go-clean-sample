# Go image for building the project
FROM golang:alpine as builder

ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

RUN mkdir -p $GOPATH/github.com/mubaiedj/go-clean-sample
WORKDIR $GOPATH/github.com/mubaiedj/go-clean-sample

COPY . .
RUN go mod vendor
COPY . .

COPY dev.config.yaml /
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOBIN/main ./app/main.go

# Runtime image with scratch container
FROM alpine
ARG VERSION
ENV VERSION_APP=$VERSION

COPY --from=builder /go/bin/ /app/
COPY --from=builder /dev.config.yaml /

ENTRYPOINT ["/app/main"]
