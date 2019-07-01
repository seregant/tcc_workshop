FROM golang:1-alpine
RUN apk add --no-cache git
RUN mkdir $GOPATH/github.com/ && mkdir $GOPATH/github.com/tugas-workshop
COPY . $GOPATH/src/github.com/tugas-workshop/
RUN go get github.com/kardianos/govendor
WORKDIR $GOPATH/src/github.com/tugas-workshop
RUN ls -al $GOPATH/github.com/tugas-workshop
RUN govendor sync
CMD ["go","run","main.go"]
EXPOSE 8000