FROM golang
WORKDIR /sub
COPY . /sub
RUN go mod download
RUN go get -d -v ./...
RUN go install -v ./...
RUN go mod tidy
EXPOSE 8081
CMD ["go","run","room.go"]