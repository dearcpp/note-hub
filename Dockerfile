FROM golang:latest
WORKDIR /usr/src/api/

COPY . .
RUN go build -o build/note-hub src/main.go

EXPOSE 8080
CMD [ "./build/note-hub" ]