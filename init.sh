#!/bin/bash

# create folder
mkdir -p data/kibana
mkdir -p data/elasticsearch
mkdir -p data/mysql
mkdir -p data/redis
mkdir -p data/rabbitmq
mkdir -p data/etcd
mkdir -p data/minio
chmod 770 data/elasticsearch && chmod 770 data/kibana
# 770 077