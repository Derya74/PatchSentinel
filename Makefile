# PatchSentinel — build, test and demo automation.
#
# Requires: Go >= 1.24, Node >= 20. No other tools are needed.

GO       ?= go
NODE     ?= node
NPM      ?= npm
BINDIR   := bin
BIN      := $(BINDIR)/patchsentinel
