FROM alpine:3.14
ENV TZ=Europe/Moscow workdir=/usr/local/bin

ADD jiraClone ${workdir}/jiraClone
CMD .${workdir}/jiraClone
