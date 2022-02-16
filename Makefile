PORT ?= COM3

flash:
	tinygo flash -target arduino -port ${PORT} main.go