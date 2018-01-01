FROM golang:1.9 AS gobuild

WORKDIR /go/src/github.com/frankh/rai-vanity
RUN go get github.com/urfave/cli \
  github.com/frankh/crypto/ed25519 \
  github.com/golang/crypto/blake2b \
  github.com/frankh/rai
COPY ./*.go ./

RUN go build -o rai-vanity .

FROM alpine

COPY --from=gobuild /go/src/github.com/frankh/rai-vanity/rai-vanity /rai-vanity

ENTRYPOINT ["/rai-vanity"]
