#!/bin/bash
CURRENT_DIR=$1
if [ -f "${CURRENT_DIR}"/.env ]; then
    echo Hello
    cat "${CURRENT_DIR}"/.env
    set -a &&. "${CURRENT_DIR}"/.env && set +a
fi