FROM golang:1.14.6-alpine3.12

RUN apk add --update make git
ADD ./ ./src/github.com/makarski/mybanana
WORKDIR ./src/github.com/makarski/mybanana

# Install gcc (needed for goose)
RUN apk add --no-cache gcc musl-dev
# Migrations
RUN go get -u github.com/pressly/goose/cmd/goose

# Application
RUN go get -u github.com/Masterminds/glide/...
RUN make build-app
EXPOSE 80
CMD make run-app
