DIR = $(shell pwd)#pwd:获得当前路径
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
RPC = $(DIR)/cmd
API_PATH= $(DIR)/cmd/api
SHELL=/bin/bash
KITEX_GEN_PATH=$(DIR)/kitex_gen
MODULE= bibi
OUTPUT=$(DIR)/output

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

SERVICES := api user video interaction follow chat
service = $(word 1, $@)
.PHONY: ${SERVICES}
$(SERVICES):
	go run $(RPC)/$(service)



.PHONY: start-all
start-all:
	sh start.sh

SERVICES := api user video interaction follow chat
.PHONY: build-all
build-all:
	@for service in $(SERVICES); do \
  		cd ${RPC};cd $$service; \
  		echo "build $$service ..." && sh build.sh; \
  		cd ${RPC}/$$service/output/bin/ && cp -r . ${OUTPUT}/$$service; \
  		echo "done"; \
  	done \

KSERVICES := user video interaction follow chat
.PHONY: kgen
kgen:
	@for kservice in $(KSERVICES); do \
		kitex -module ${MODULE} ${IDL_PATH}/$$kservice.thrift; \
    	cd ${RPC};cd $$kservice;kitex -module ${MODULE} -service $$kservice -use ${KITEX_GEN_PATH} ${IDL_PATH}/$$kservice.thrift; \
    	cd ../../; \
    done \


.PHONY: hzgen
hzgen:
	cd ${API_PATH}; \
	hz update -idl ${IDL_PATH}/api.thrift; \
	swag init; \

