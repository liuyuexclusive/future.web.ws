FROM alpine
ADD future.web.ws /future.web.ws
ADD appconfig.yml /appconfig.yml
ENTRYPOINT [ "/future.web.ws" ]
