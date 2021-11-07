FROM hashicorp/terraform:light

RUN apk --no-cache --update add tzdata dumb-init

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod 700 /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["dumb-init", "--", "docker-entrypoint.sh"]
CMD ["provision"]