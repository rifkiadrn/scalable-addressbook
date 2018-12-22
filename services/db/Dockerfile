FROM postgres:10.3
COPY sql/init.sql /docker-entrypoint-initdb.d/0.sql
COPY sql/up.sql /docker-entrypoint-initdb.d/1.sql

RUN chmod -R 775 /docker-entrypoint-initdb.d

CMD ["postgres"]