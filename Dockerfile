FROM golang:1.14.0-alpine

COPY . ./bitbucket.org/tv2norge/sudo-gometrics-demo

WORKDIR $GOPATH/bitbucket.org/tv2norge/sudo-gometrics-demo

RUN go build .

RUN ls -la
RUN echo "Jadda"


CMD ["./sudo-gometrics-demo"]