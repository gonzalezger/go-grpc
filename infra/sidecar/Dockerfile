FROM envoyproxy/envoy-alpine-dev:latest

COPY envoy.yaml.tmpl /tmpl/envoy.yaml.tmpl
COPY docker-entrypoint.sh /

RUN chmod 500 /docker-entrypoint.sh

RUN apk update && apk add gettext

ENV LISTEN_PORT=90 SERVICE_ADDRESS=service SERVICE_PORT=9090

EXPOSE 90

ENTRYPOINT ["/docker-entrypoint.sh"]