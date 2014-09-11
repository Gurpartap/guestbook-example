FROM google/golang

RUN go get github.com/Gurpartap/guestbook-example && \
    cp /gopath/src/github.com/Gurpartap/guestbook-example/minimal-image/Dockerfile /gopath

CMD docker build --rm --force-rm -t gurpartap/guestbook-example /gopath
