FROM golang:1.21 AS build-stage
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /sava .

FROM scratch
COPY --from=build-stage /sava /sava
ENTRYPOINT ["/sava"]
