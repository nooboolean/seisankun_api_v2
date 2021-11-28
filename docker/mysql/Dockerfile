FROM mysql:8.0.21

RUN chown -R mysql /var/lib/mysql && \
    chgrp -R mysql /var/lib/mysql

COPY ./my.conf /etc/mysql/conf.d/my.conf

CMD ["mysqld"]

EXPOSE 3306