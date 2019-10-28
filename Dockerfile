FROM alpine
ADD future.web.ws /future.web.ws
ADD appconfig.json /appconfig.json
ENTRYPOINT [ "/future.web.ws" ]
