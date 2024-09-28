FROM debian:stable

COPY ./.bin/api ./bin/api
COPY ./migrations ./migrations

CMD [ "/bin/api" ]