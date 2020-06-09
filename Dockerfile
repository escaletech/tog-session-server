FROM alpine

WORKDIR /app

COPY ./dist/server /app/server

EXPOSE 3000

CMD ["/app/server"]
