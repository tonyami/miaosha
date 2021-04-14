FROM alpine

ADD bin/miaosha /

RUN chmod 777 /miaosha

ENTRYPOINT ["sh", "-c", "/miaosha"]
