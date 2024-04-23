#!/bin/bash

(go run ./cmd/api) &
sleep 2s
printf "\n"

(go run ./cmd/user) &
sleep 2s
printf "\n"

(go run ./cmd/video) &
sleep 2s
printf "\n"

(go run ./cmd/interaction) &

wait