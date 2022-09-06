FROM bash:5.1

COPY ./cmd/usesthisreader/usesthisreader /usesthisreader

COPY entrypoint.sh entrypoint.sh

RUN chmod a+x /entrypoint.sh

CMD ["./entrypoint.sh"]
