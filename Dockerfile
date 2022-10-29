FROM golang:alpine

WORKDIR /task6-api/api/

COPY . .

ENV DB_HOST=${DB_HOST}
ENV DB_USER=${DB_USER}
ENV DB_PASS=${DB_PASS}
ENV DB_NAME=${DB_NAME}
ENV APP_PORT = ${APP_PORT}

RUN go mod tidy
RUN go build -o start
RUN /task6-api/api/start migrate -u

CMD [ "/task6-api/api/start serve" ]