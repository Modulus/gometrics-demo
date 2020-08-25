FROM golang:1.15.0-alpine

COPY . ./bitbucket.org/tv2norge/sudo-gometrics-demo

WORKDIR $GOPATH/bitbucket.org/tv2norge/sudo-gometrics-demo

RUN go build .

RUN ls -la
RUN echo "Jadda"

EXPOSE 8000
EXPOSE 2112

CMD ["./sudo-gometrics-demo"]