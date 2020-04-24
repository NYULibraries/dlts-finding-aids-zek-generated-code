#!/bin/bash

git clone git@github.com:NYULibraries/findingaids_eads.git

find findingaids_eads/ -type f -name *.xml -exec cat {} \; | zek | grep -v '// Ead was generated' > published-finding-aids-ead-files.go

rm -fr findingaids_eads/

git clone git@github.com:NYULibraries/dlts-finding-aids-ead-sample-set-1.git

find dlts-finding-aids-ead-sample-set-1/ -type f -name *.xml -exec cat {} \; | zek | grep -v '// Ead was generated' > sample-set-1-ead-files.go

rm -fr dlts-finding-aids-ead-sample-set-1/

