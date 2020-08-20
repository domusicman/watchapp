FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
# had to make a few changes since the repo has been updated.
RUN go get -u github.com/go-sql-driver/mysql