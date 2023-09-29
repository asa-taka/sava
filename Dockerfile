FROM golang:1.21 AS build-stage
WORKDIR /src
COPY . .
RUN go build -o /sava .

FROM scratch
COPY --from=build-stage /sava /sava
ENTRYPOINT ["/sava"]
