FROM debian:latest

RUN apt -y update
RUN apt -y install fortune-mod fortunes fortunes-off openssl ca-cacert

COPY dirtyrobocop /bin/
COPY dirtyrobocop.json /etc/

CMD ["/bin/dirtyrobocop", "/etc/dirtyrobocop.json", "runforever"]
