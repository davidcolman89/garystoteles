FROM golang:1.10
ENV GOBIN /go/bin

WORKDIR /go/src/github.com/davidcolman89/garystoteles
ADD src/api .

# Go dep!
#RUN go get -u github.com/golang/dep/...
#RUN dep ensure

# Build my app
CMD ["go","run","main.go"]

EXPOSE 9090
