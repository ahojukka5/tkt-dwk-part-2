#!/bin/sh

set -e

docker build -t ahojukka5/dwk-mainapp-gen-timestamp mainapp/gen-timestamp
docker push ahojukka5/dwk-mainapp-gen-timestamp

docker build -t ahojukka5/dwk-mainapp-read-timestamp mainapp/read-timestamp
docker push ahojukka5/dwk-mainapp-read-timestamp

docker build -t ahojukka5/dwk-pingpong pingpong
docker push ahojukka5/dwk-pingpong

docker build -t ahojukka5/dwk-todo-backend todo-backend
docker push ahojukka5/dwk-todo-backend

docker build -t ahojukka5/dwk-todo-frontend todo-frontend
docker push ahojukka5/dwk-todo-frontend

docker build -t ahojukka5/dwk-todo-cronjob todo-cronjob && docker push ahojukka5/dwk-todo-cronjob

docker build -t ahojukka5/dwk-secrets secrets
docker push ahojukka5/dwk-secrets
