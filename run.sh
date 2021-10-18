#!/bin/sh
docker rm delete
docker rmi -f $(docker images 'ariorh/jira_clone:v1' -a -q)
docker build -t ariorh/jira_clone:v1 . &&
docker run --name delete ariorh/jira_clone:v1