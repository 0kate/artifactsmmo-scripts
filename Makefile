CMD_DIR := ./cmd
OUTPUT_DIR := ./out

all: apps artifacts-cli

monsters: find-all-monsters

apps: monster-buster

find-all-monsters:
	go build -o $(OUTPUT_DIR)/find-all-monsters $(CMD_DIR)/find-all-monsters/main.go

monster-buster:
	go build -o $(OUTPUT_DIR)/monster-buster $(CMD_DIR)/monster-buster/main.go

artifacts-cli:
	go build -o $(OUTPUT_DIR)/artifacts-cli $(CMD_DIR)/artifacts-cli/main.go
