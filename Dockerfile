FROM golang:1.10 AS gobuild

WORKDIR /go/src/github.com/frankh/nano-vanity
RUN go get github.com/urfave/cli \
  github.com/frankh/crypto/ed25519 \
  github.com/golang/crypto/blake2b \
  github.com/frankh/nano
COPY ./*.go ./

RUN go build -o nano-vanity .

FROM alpine

COPY --from=gobuild /go/src/github.com/frankh/nano-vanity/nano-vanity /nano-vanity

ENTRYPOINT ["/nano-vanity"]
