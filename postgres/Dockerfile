FROM postgres:13


RUN whoami
ADD db/01-init.sh /docker-entrypoint-initdb.d/
# wtf windows...
RUN sed -i -e 's/^M$//' /docker-entrypoint-initdb.d/01-init.sh
RUN sed -i -e 's/\r$//' /docker-entrypoint-initdb.d/01-init.sh
ENTRYPOINT ["docker-entrypoint.sh"]
EXPOSE 5432
CMD ["postgres"]