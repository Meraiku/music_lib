FROM debian:stable

COPY ./.bin/api ./bin/api

CMD [ "/bin/api" ]