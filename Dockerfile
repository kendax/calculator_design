FROM golang:1.16.4-buster

RUN mkdir -p /calculator_design

WORKDIR /calculator_design

ADD . /calculator_design

RUN go build -o main_program main_program.go

CMD ["./main_program"]