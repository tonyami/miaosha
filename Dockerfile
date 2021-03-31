FROM alpine

ADD bin/miaosha /

ADD conf.ini /

RUN chmod 777 /miaosha

ENTRYPOINT ["sh", "-c", "/miaosha"]
