FROM ubuntu:18.04

RUN apk update && apk upgrade

RUN apk install libmagickwand-dev

COPY . .

RUN cp .env.yaml.example .env.yaml

RUN make

EXPOSE 8080

CMD ["app", "stage=dev"]