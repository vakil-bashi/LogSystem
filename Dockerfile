FROM golang:1.20
EXPOSE 9091

WORKDIR /prj

COPY . /prj/
RUN apt-get install gcc
RUN go build main.go
CMD ./main