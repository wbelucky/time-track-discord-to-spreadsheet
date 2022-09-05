FROM golang:1.18-alpine as build-env
ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}

# ref: https://github.com/GoogleContainerTools/distroless/blob/main/examples/go/Dockerfile
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app cmd/main.go


FROM gcr.io/distroless/static

COPY --from=build-env /go/bin/app /
ENV GO_ENV=prd
CMD ["/app"]

