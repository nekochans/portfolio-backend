FROM mysql:8.0.23

LABEL maintainer="https://github.com/nekochans"

COPY ./docker/mysql/config/my.cnf /etc/mysql/conf.d/my.cnf

RUN set -eux && \
  mkdir /var/log/mysql && \
  chown mysql:mysql /var/log/mysql
