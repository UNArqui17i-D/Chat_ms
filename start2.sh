#!/usr/bin/env sh

mvn -f pom.xml package

./../rancher-compose --project-name Chat-ms \
    --url http://192.168.99.100:8080/v1/projects/1a5 \
    --access-key A488D9A45F7630739A2D \
    --secret-key i4X428DBgAtakraeiQfrFCeFnogt6oaUnHbBf6PS \
    -f docker-compose.yml \
    --verbose up \
    -d --force-upgrade \
    --confirm-upgrade
