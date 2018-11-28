FROM golang

ADD . /logger
WORKDIR /logger
RUN go build 

ENTRYPOINT ["./logger"]