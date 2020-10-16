FROM envoyproxy/envoy-alpine-dev:latest

COPY envoy.yaml.tmpl /tmpl/envoy.yaml.tmpl
COPY docker-entrypoint.sh /

RUN chmod 500 /docker-entrypoint.sh

RUN apk update && apk add gettext

ENV LISTEN_PORT=90 SERVICE_DISCOVERY_ADDRESS=service SERVICE_DISCOVERY_PORT=${SERVICE_DISCOVERY_PORT}

ENTRYPOINT ["/docker-entrypoint.sh"]