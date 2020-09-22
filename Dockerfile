FROM scratch
LABEL MAINTAINER="Aleksandr Glukhov"

ADD simple_server /

EXPOSE 8080
CMD ["/simple_server"]
