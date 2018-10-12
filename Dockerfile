FROM golang:1.11

EXPOSE 8080

RUN apt-get update
RUN apt-get install -y apt-utils
RUN apt-get install -y sqlite3
RUN apt-get install -y make
WORKDIR /go/src/starter-x-golang
COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...
RUN make install

CMD /go/bin/main