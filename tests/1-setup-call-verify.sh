#!/bin/sh
set -e

MOCKED_EXE=/usr/local/bin/mongod

echo "> testing happy path"
shellmock setup ${MOCKED_EXE}
${MOCKED_EXE}
${MOCKED_EXE} arg1 arg2
shellmock verify ${MOCKED_EXE}
shellmock verify ${MOCKED_EXE} arg1 arg2
shellmock verifyNoInteraction ${MOCKED_EXE}

echo "> testing calls and failing verify because of different arg strings"
shellmock setup ${MOCKED_EXE}
${MOCKED_EXE}
${MOCKED_EXE} arg1 arg2
shellmock verify ${MOCKED_EXE}
shellmock verify ${MOCKED_EXE} arg1 arg3 || echo "expected error. i am fine"

echo "> testing calls and failing verify because of different arg counts"
shellmock setup ${MOCKED_EXE}
${MOCKED_EXE}
${MOCKED_EXE} arg1 arg2
shellmock verify ${MOCKED_EXE}
shellmock verify ${MOCKED_EXE} arg1 || echo "expected error. i am fine"

echo "> testing calls and failing verify because of different call counts"
shellmock setup ${MOCKED_EXE}
${MOCKED_EXE}
${MOCKED_EXE} arg1 arg2
shellmock verify ${MOCKED_EXE}
shellmock verify ${MOCKED_EXE} arg1 arg2
shellmock verify ${MOCKED_EXE} arg3 arg4 || echo "expected error. i am fine"

echo "> testing calls and failing verifyNoInteraction"
shellmock setup ${MOCKED_EXE}
${MOCKED_EXE}
${MOCKED_EXE} arg1 arg2
shellmock verify ${MOCKED_EXE}
shellmock verifyNoInteraction ${MOCKED_EXE} || echo "expected error. i am fine"

echo "> testing no call and failing verify"
shellmock setup ${MOCKED_EXE}
shellmock verify ${MOCKED_EXE} || echo "expected error. i am fine"

echo "> testing no call and verifyNoInteraction"
shellmock setup ${MOCKED_EXE}
shellmock verifyNoInteraction ${MOCKED_EXE}
