# https://hub.docker.com/_/scratch
FROM scratch

# Copy the web binary to the container.
COPY bin/web /

# Run the web binary.
CMD ["/web"]