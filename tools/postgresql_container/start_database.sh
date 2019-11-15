#!/bin/sh

docker run --name postgresql\
	-v /home/vagrant/tools/postgresql_container/postgresql:/mnt\
	-e POSTGRES_USER=sample\
	-e POSTGRES_PASSWORD=sample\
	-e POSTGRES_DB=sampledb\
	-p 5432:5432\
	-d postgres:11.3
