FROM node:16.15.0-alpine as client_builder

WORKDIR /app
COPY client/package*.json ./

RUN npm install

ADD client /app/

RUN npm run build

FROM golang:1.18.2-alpine as server_builder

WORKDIR /app

COPY server/go.mod ./
COPY server/go.sum ./
RUN go mod download

RUN go mod tidy

ADD server /app/

RUN go build -o /the-archives

FROM nginx:alpine

COPY --from=client_builder /app/public /usr/share/nginx/html
COPY --from=server_builder /the-archives /usr/bin/the-archives

COPY .nginx/default.conf /etc/nginx/conf.d/

CMD ["/usr/bin/the-archives"]