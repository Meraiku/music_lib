FROM debian:stable

COPY ./.bin/api ./bin/api
COPY ./sql/migrations ./sql/migrations

CMD [ "/bin/api" ]