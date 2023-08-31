FROM postgres:14-alpine3.15
RUN mkdir ./db
COPY . db/
RUN chown -R root ./db

WORKDIR /db

RUN chmod +x createDB.sh