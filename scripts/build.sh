#!/bin/bash
cd ..
go build .
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
echo "Darwin built into ../Darwin. Run install.sh to deploy that to your /usr/local/bin."
