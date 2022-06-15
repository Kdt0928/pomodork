# 使うDockerイメージ
FROM --platform=linux/x86_64 ruby:3.1.2

ENV RIDGEPOLE_VERSION 1.0.7
ENV MYSQL_VERSION 0.5.4

RUN gem install ridgepole --version $RIDGEPOLE_VERSION \
    && gem install mysql2 --version $MYSQL_VERSION

WORKDIR /ridgepole
ADD ./config/database.yml ./config/database.yml
ADD ./schemas/Schemafile ./schemas/Schemafile