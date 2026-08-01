#!/bin/bash
set -euxo pipefail

# do the release.
if [[ $GITHUB_REF == refs/tags/v* ]]; then
  make release
else
  make release-snapshot
fi
