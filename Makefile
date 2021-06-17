CURRENT_DIR=$(shell pwd)

set-env:
	./shell_scripts/set-env.sh ${CURRENT_DIR}

gen-swag:
	swag init -g api/router.go -o api/docs