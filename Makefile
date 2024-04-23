DIR = $(shell pwd)#pwd:获得当前路径
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
RPC = $(DIR)/cmd
API_PATH= $(DIR)/cmd/api
SHELL=/bin/bash
MODULE= bibi

.PHONY: init
init:
	mv config/config-example.yaml config/config.yaml
	sh init.sh

.PHONY: env-up
env-up:
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down

SERVICES := api user video interaction
service = $(word 1, $@)
.PHONY: $(SERVICES)
$(SERVICES):
	go run $(RPC)/$(service)



.PHONY: build-all
build-all:
	sh start.sh


KSERVICES := user video interaction
.PHONY: kgen
kgen:
	@for kservice in $(KSERVICES); do \
#  		sh kitex_update.sh $$kservice; \
		kitex -module ${MODULE} idl/$$kservice.thrift; \
    	cd ${RPC};cd $$kservice;kitex -module ${MODULE} -service $$kservice -use bibi/kitex_gen ../../idl/$$kservice.thrift; \
    done \
    echo "done"

.PHONY: hzgen
hzgen:
	cd ${API_PATH}; \
	hz update -idl ${IDL_PATH}/api.thrift; \

