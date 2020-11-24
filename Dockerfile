FROM golang:latest
WORKDIR /usr/src/note-hub/

COPY . .
RUN go build -o build/note-hub cmd/main.go

EXPOSE 8080
CMD [ "./build/note-hub" ]