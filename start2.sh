#!/usr/bin/env sh

./../rancher-compose --project-name chat-ms \
    --url http://192.168.99.100:8080/v1/projects/1a5 \
    --access-key 187A10CF0B968A0BC159 \
    --secret-key Vy68nzh4oo1raWEjqHhCHRGzcUfqqPHT7NUXQcMw \
    -f docker-compose.yml \
    --verbose up \
    -d --force-upgrade \
    --confirm-upgrade