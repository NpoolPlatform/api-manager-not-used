FROM jrei/systemd-centos:7

COPY ApiManager.viper.yaml /etc/ApiManager/ApiManager.viper.yaml
COPY api-manager /usr/local/bin/api-manager

CMD ["api-manager", "run"]
