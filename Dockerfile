FROM debian:latest

RUN apt -y update
RUN apt -y install fortune-mod fortunes fortunes-off

COPY dirtyrobocop /bin/
COPY config.json /etc/

CMD ["/bin/dirtyrobocop", "/etc/config.json"]
