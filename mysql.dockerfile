FROM mysql:5.7

RUN apt-get update \
    && apt-get install -y --no-install-recommends tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && mysql_tzinfo_to_sql /usr/share/zoneinfo/Asia/Tokyo 'Asia/Tokyo' \
    && rm -rf /var/cache/apt/*