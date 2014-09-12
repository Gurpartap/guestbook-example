FROM busybox:ubuntu-14.04

ADD ./bin/guestbook-example /bin/guestbook-example
ADD ./src/github.com/Gurpartap/guestbook-example/public/index.html /public/index.html
ADD ./src/github.com/Gurpartap/guestbook-example/public/script.js /public/script.js
ADD ./src/github.com/Gurpartap/guestbook-example/public/style.css /public/style.css

EXPOSE 3000
CMD ["/bin/guestbook-example"]
