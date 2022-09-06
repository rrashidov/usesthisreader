FROM ubuntu:latest
RUN useradd usesthisreader

FROM alpine:3.16.2

COPY ./bin/linux_arm/usesthisreader /home/usesthisreader/usesthisreader
COPY ./docker/entrypoint.sh /home/usesthisreader/entrypoint.sh
RUN chmod a+x /home/usesthisreader/entrypoint.sh

COPY --from=0 /etc/passwd /etc/passwd

USER usesthisreader

WORKDIR /home/usesthisreader/

CMD ["/home/usesthisreader/entrypoint.sh"]
