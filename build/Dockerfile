FROM debian:jessie
WORKDIR /opt
EXPOSE 8080

RUN chmod 755 /usr/share/easy-rsa/*
ADD assets/start.sh /opt/start.sh

ADD openvpn-manage.tar.gz /opt/openvpn-gui/
RUN rm -f /opt/openvpn-gui/data.db
ADD assets/app.conf /opt/openvpn-gui/conf/app.conf

CMD /opt/start.sh
