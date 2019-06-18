FROM node:lts-alpine AS build

WORKDIR /build

ENV VUE_APP_BASE_URL http://192.168.1.3:3000
# ENV VUE_APP_NGINX_PROXY /stream-app
ENV NODE_ENV production

RUN apk --no-cache add git 
RUN git clone https://github.com/stream-app-go/frontend.git . \
&& yarn && yarn run build

FROM golang:1.12-alpine

ENV GO111MODULE=on
ENV MONGO mongodb://mongo:27018

WORKDIR $GOPATH/src/github.com/HichuYamichu/stream-app-server
RUN apk --no-cache add git ffmpeg

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
COPY --from=build /build/dist ./web/dist

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./stream-app-server ./cmd/main.go

CMD ./stream-app-server