#!/usr/bin/env bash
# Install pre-commit and required dependencies.
set -e

echo "--Installing pre-commit & dependencies---"
brew install pre-commit
brew install golangci-lint
pre-commit install


# shellcheck disable=SC2181
if [ "$?" = "0" ]; then
    echo "--- setup successful ---"
else
    echo "--- setup failed ---"
fi