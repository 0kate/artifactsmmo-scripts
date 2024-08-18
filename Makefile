CMD_DIR := ./cmd
OUTPUT_DIR := ./out

all: introduction maps apps

introduction: move fight gathering unequip crafting equip

maps: find-map find-all-maps

monsters: find-all-monsters

apps: monster-buster

move:
	go build -o $(OUTPUT_DIR)/move $(CMD_DIR)/move/main.go

fight:
	go build -o $(OUTPUT_DIR)/fight $(CMD_DIR)/fight/main.go

gathering:
	go build -o $(OUTPUT_DIR)/gathering $(CMD_DIR)/gathering/main.go

unequip:
	go build -o $(OUTPUT_DIR)/unequip $(CMD_DIR)/unequip/main.go

crafting:
	go build -o $(OUTPUT_DIR)/crafting $(CMD_DIR)/crafting/main.go

equip:
	go build -o $(OUTPUT_DIR)/equip $(CMD_DIR)/equip/main.go

find-map:
	go build -o $(OUTPUT_DIR)/find-map $(CMD_DIR)/find-map/main.go

find-all-maps:
	go build -o $(OUTPUT_DIR)/find-all-maps $(CMD_DIR)/find-all-maps/main.go

find-all-monsters:
	go build -o $(OUTPUT_DIR)/find-all-monsters $(CMD_DIR)/find-all-monsters/main.go

monster-buster:
	go build -o $(OUTPUT_DIR)/monster-buster $(CMD_DIR)/monster-buster/main.go
