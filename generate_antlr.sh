#!/bin/sh
set -eu

# Generates the lexer and the parser from the grammar into ./pkg/a2l/parser/grammar.
#
# The generator is pinned: the code it emits must match the runtime the module depends on
# (github.com/antlr4-go/antlr/v4 in go.mod), so a floating version would silently produce sources
# which no longer fit it, or change the parser without a change in this repository. Keep both in
# step when either of them is updated.
ANTLR_VERSION=4.13.2
ANTLR_SHA256=eae2dfa119a64327444672aff63e9ec35a20180dc5b8090b7a6ab85125df4d76
ANTLR_JAR="antlr-${ANTLR_VERSION}-complete.jar"

curl -fsSLo "${ANTLR_JAR}" "https://www.antlr.org/download/${ANTLR_JAR}"
echo "${ANTLR_SHA256}  ${ANTLR_JAR}" | sha256sum -c -

java -jar "${ANTLR_JAR}" -Dlanguage=Go -visitor -o ./pkg/a2l/parser ./grammar/A2L.g4

rm -f "${ANTLR_JAR}"
