FROM golang:1.8 as backend
RUN mkdir -p /go/src/github.com/frederikleemans/e3w
ADD . /go/src/github.com/frederikleemans/e3w
WORKDIR /go/src/github.com/frederikleemans/e3w
RUN CGO_ENABLED=0 go get -d -v ./...
RUN CGO_ENABLED=0 go build -i

FROM node:8 as frontend
RUN mkdir /app
ADD static /app
WORKDIR /app
RUN npm --registry=https://registry.npm.taobao.org \
--cache=$HOME/.npm/.cache/cnpm \
--disturl=https://npm.taobao.org/mirrors/node \
--userconfig=$HOME/.cnpmrc install && npm run publish

FROM alpine:latest
RUN mkdir -p /app/static/dist /app/conf
COPY --from=backend /go/src/github.com/frederikleemans/e3w/e3w /app
COPY --from=frontend /app/dist /app/static/dist
COPY conf/config.default.ini /app/conf

# Create user HOME
RUN mkdir /app/home

# Create a user group 'e3wg'
RUN addgroup -S e3wg

# Create a user 'e3wu' under 'e3wg'
RUN adduser -S -D -h /app/home e3wu e3wg

# Chown all the files to the app user.
RUN chown -R e3wu:e3wg /app

# Switch to 'e3wu'
USER e3wu

EXPOSE 8080
WORKDIR /app
CMD ["./e3w"]
