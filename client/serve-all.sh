#!/bin/sh
ng serve &

cd  ../server/

go run main.go &

wait