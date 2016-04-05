#!/bin/bash

set -e

pushd ttn > /dev/null
echo 'Updating ttn cli docs...'
HOME='$HOME' go run update.go
popd > /dev/null

pushd ttnctl > /dev/null
echo 'Updating ttnctl cli docs...'
HOME='$HOME' go run update.go
popd > /dev/null

echo 'Preparing git commit...'
git add ttn/ttn.md ttn/ttn_*.md ttnctl/ttnctl.md ttnctl/ttnctl_*.md

echo "To commit, run: git commit -m 'Update cli docs'"
