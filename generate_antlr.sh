#!/bin/sh
set -eu

# Generates the lexer and the parser from the grammar into ./pkg/a2l/parser/grammar.
#
# The generator is pinned: the code it emits must match the runtime the module depends on
# (github.com/antlr4-go/antlr/v4 in go.mod), so a floating version would silently produce sources
# which no longer fit it, or change the parser without a change in this repository. Keep both in
# step when either of them is updated.
ANTLR_VERSION=4.13.1
ANTLR_SHA256=bc13a9c57a8dd7d5196888211e5ede657cb64a3ce968608697e4f668251a8487
ANTLR_JAR="antlr-${ANTLR_VERSION}-complete.jar"

curl -fsSLo "${ANTLR_JAR}" "https://www.antlr.org/download/${ANTLR_JAR}"
echo "${ANTLR_SHA256}  ${ANTLR_JAR}" | sha256sum -c -

java -jar "${ANTLR_JAR}" -Dlanguage=Go -visitor -o ./pkg/a2l/parser ./grammar/A2L.g4

rm -f "${ANTLR_JAR}"
