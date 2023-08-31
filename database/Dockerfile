FROM alpine:latest

WORKDIR /database

ADD *.sql database/
ADD createDB.sh .

RUN chmod +x createDB.sh

ENTRYPOINT ["bash", "createDB.sh"]