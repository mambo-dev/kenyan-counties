FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD kenyan_locations /usr/bin/kenyan_locations

CMD ["kenyan_locations"]
