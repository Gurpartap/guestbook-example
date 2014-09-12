FROM busybox:ubuntu-14.04

ADD ./bin/guestbook-example /app/guestbook-example
ADD ./src/github.com/Gurpartap/guestbook-example/public/index.html /app/index.html
ADD ./src/github.com/Gurpartap/guestbook-example/public/script.js /app/script.js
ADD ./src/github.com/Gurpartap/guestbook-example/public/style.css /app/style.css

EXPOSE 3000
CMD ["/app/guestbook-example"]
