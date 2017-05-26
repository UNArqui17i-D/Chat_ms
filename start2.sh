#!/usr/bin/env sh

./../rancher-compose --project-name chat-ms \
    --url http://192.168.99.100:8080/v1/projects/1a5 \
    --access-key D0C832691E321859CD8D \
    --secret-key fjN1tMEDZRYHcB1zEuSsoKVmF7NtPahNWX3uBWE2 \
    -f docker-compose.yml \
    --verbose up \
    -d --force-upgrade \
    --confirm-upgrade