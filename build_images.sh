#!/usr/bin/sh

docker build -t ahojukka5/dwk-todo-backend todo-backend
docker push ahojukka5/dwk-todo-backend

docker build -t ahojukka5/dwk-todo-frontend todo-frontend
docker push ahojukka5/dwk-todo-frontend
