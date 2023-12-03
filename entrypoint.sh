#! /bin/sh

ARGS=""

if [ -n "${SOLAX_TOKEN_ID}" ]; then
    ARGS="${ARGS} --token-id ${SOLAX_TOKEN_ID}"
fi

if [ -n "${SOLAX_SN}" ]; then
    ARGS="${ARGS} --sn ${SOLAX_SN}"
fi

./solax_exporter ${ARGS} $@
