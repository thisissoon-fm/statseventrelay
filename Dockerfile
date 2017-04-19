# Lightweight Hardended Linux Distro
FROM alpine:3.5

# Update and Install OS level packages
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Default build arguments
ARG BINLOC=./statseventrelay.linux-amd64
ARG BINDEST=/usr/local/bin/statseventrelay

# Add Crest user
RUN adduser -D -H soon_

# Copy Binary
COPY ${BINLOC} ${BINDEST}

# Volumes
VOLUME ["/etc/sfm/statseventrelay", "/var/log/sfm/statseventrelay"]

# Set our Application Entrypoint
ENTRYPOINT ["statseventrelay"]
