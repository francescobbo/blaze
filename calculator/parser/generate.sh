#!/bin/bash

set -euo pipefail

shopt -s expand_aliases

alias antlr4='java -Xmx500M -cp "../../tools/antlr-4.13.1-complete.jar:\$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -no-visitor -package parser Calculator.g4