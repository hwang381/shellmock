#!/bin/sh
set -e

MOCKED_EXE=/usr/local/bin/mongod
MOCKED_EXE_STATE=/tmp/usr_local_bin_mongod.json

shellmock setup ${MOCKED_EXE}
${MOCKED_EXE} arg0
shellmock setup ${MOCKED_EXE}
if test -f "${MOCKED_EXE_STATE}"; then
  echo "${MOCKED_EXE_STATE} exists"
  exit 1
fi
