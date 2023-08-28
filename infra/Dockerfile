# adapted from here https://github.com/confluentinc/confluent-kafka-go/issues/461
FROM golang:1.20-alpine AS builder

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo

COPY . .

RUN go build -tags musl --ldflags "-extldflags -static" -o app .

FROM scratch AS runner

COPY --from=builder /go/src/app /

EXPOSE 8080

ENTRYPOINT ["./app"]
