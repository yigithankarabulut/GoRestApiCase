FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=roots-secret-pw
ENV MYSQL_DATABASE=my_database

ENV MYSQL_CHARSET=utf8
ENV MYSQL_COLLATION=utf8_general_ci

CMD ["mysqld"]

EXPOSE 3306
