FROM jrei/systemd-centos:7

COPY APIManager.viper.yaml /etc/APIManager/APIManager.viper.yaml
COPY api-manager /usr/local/bin/api-manager

CMD ["api-manager", "run"]
