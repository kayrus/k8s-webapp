FROM alpine:3.3
# go get github.com/josephspurrier/gowebapp
# CGO_ENABLED=0 go build -a -installsuffix cgo
WORKDIR /app

ADD static /app/static
ADD config /app/config
ADD template /app/template
ADD entrypoint.sh gowebapp /app/

EXPOSE 8080
CMD /app/entrypoint.sh
