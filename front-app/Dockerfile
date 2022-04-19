FROM node:12.4.0-alpine AS buildvue
WORKDIR /app
COPY front /app
RUN npm install -g @vue/cli
RUN npm install --silent && npm run build
RUN ls -all /app
# production environment
FROM golang AS gobuilder
ENV Version="2.0.0" 
WORKDIR /go/src/github.com/jpgouin/
COPY back/api/ .
#ENV GOPATH=/go/src/github.com/jpgouin/
RUN ls -all /go/src/github.com/jpgouin/
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api *.go
RUN ls -all /go/src/github.com/jpgouin/

FROM nginx:1.16.0-alpine
ENV Version="2.0.0" 
RUN apk add --no-cache ca-certificates 
COPY --from=buildvue /app/dist /usr/share/nginx/html
COPY --from=gobuilder /go/src/github.com/jpgouin/api /usr/share/
COPY default.conf /etc/nginx/conf.d/default.conf
COPY start.sh /usr/local/bin/
EXPOSE 80
CMD ["start.sh"]