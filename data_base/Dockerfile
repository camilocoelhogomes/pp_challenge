# Derivando da imagem oficial do MySQL
FROM docker.io/postgres:latest
ENV POSTGRES_PASSWORD admin
ADD init.sql /docker-entrypoint-initdb.d
EXPOSE 5432
