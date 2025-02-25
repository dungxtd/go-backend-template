FROM golang:1.23.3-alpine as build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN go build -o /main ./

FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /main /main
COPY --from=build /app/.env /app/.env
COPY --from=build /app/templates /app/templates

WORKDIR /app
EXPOSE 8080
CMD ["/main"]