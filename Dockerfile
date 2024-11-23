FROM golang:1.23 
WORKDIR /go/src/app 
COPY . . 
EXPOSE 8080
CMD ["go", "run", "/cmd/web"]