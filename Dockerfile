FROM golang:1.12.7-alpine3.10 AS build
RUN apk add --no-cache make
WORKDIR /go/src/github.com/keilerkonzept/aws-secretsmanager-files/
COPY . .
ENV CGO_ENABLED=0
RUN make binaries/linux_x86_64/aws-secretsmanager-files && mv binaries/linux_x86_64/aws-secretsmanager-files /app

FROM alpine:3.11
RUN apk add --no-cache ca-certificates
COPY --from=build /app /bin/aws-secretsmanager-files
CMD [ "aws-secretsmanager-files" ]
