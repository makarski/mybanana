FROM golang:1.8.3-alpine

RUN apk add --update make git
ADD ./ ./src/github.com/makarski/mybanana
WORKDIR ./src/github.com/makarski/mybanana
RUN make build-app
EXPOSE 8000
CMD make run-app
