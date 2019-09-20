FROM   golang:alpine

RUN apk update

COPY . /go/src

WORKDIR /go/src

RUN go build -o div .

EXPOSE 7072

RUN chmod 755 div

CMD [ "./div" ]