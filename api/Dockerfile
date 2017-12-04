FROM golang:1.9.2

WORKDIR /go/src/git.arthrex.io/dschultz/surgeon-a/api
COPY . .

RUN go install ./vendor/github.com/ddollar/init
RUN go install ./vendor/github.com/cleblanc87/rerun
RUN go install

ENTRYPOINT ["init"]
CMD ["bash", "bin/api.sh"]
