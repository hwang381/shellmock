#!/bin/sh
set -e

MOCKED_EXE=/usr/local/bin/mongod

shellmock setup ${MOCKED_EXE}
${MOCKED_EXE} arg0
${MOCKED_EXE} arg1 arg2
shellmock verify ${MOCKED_EXE} arg0
shellmock verify ${MOCKED_EXE} arg1 arg2
shellmock verifyNoInteraction ${MOCKED_EXE}
