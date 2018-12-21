FROM postgres:10.3
COPY init.sql /docker-entrypoint-initdb.d/0.sql
COPY up.sql /docker-entrypoint-initdb.d/1.sql
CMD ["postgres"]