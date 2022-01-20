FROM golang:1.15.3
WORKDIR /app/src/
COPY . .
RUN go build -o /app.py
COPY go.mod ./
RUN go mod download
RUN go build -o app /app/src/cmd/main.go
COPY cmd/main.go go.* /src/cmd

CMD ["go", "run", "/app/src/cmd/main.go"]
