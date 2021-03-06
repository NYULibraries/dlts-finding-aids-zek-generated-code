#!/bin/bash

# Current EAD files repo
git clone git@github.com:NYULibraries/findingaids_eads.git

find findingaids_eads/ -type f -name *.xml -exec cat {} \; | zek | grep -v '// Ead was generated' > published-finding-aids-ead-files.go

rm -fr findingaids_eads/

# SAMPLE SET 1
git clone git@github.com:NYULibraries/dlts-finding-aids-ead-sample-set-1.git

find dlts-finding-aids-ead-sample-set-1/ -type f -name *.xml -exec cat {} \; | zek | grep -v '// Ead was generated' > sample-set-1-ead-files.go

rm sample-set-1-individual-ead-files/*

for eadfile in $( find dlts-finding-aids-ead-sample-set-1/ -type f -name *.xml ); do eadid=$( basename $eadfile .xml ); cat $eadfile | zek | grep -v '// Ead was generated' > sample-set-1-individual-ead-files/${eadid}.go; done

rm -fr dlts-finding-aids-ead-sample-set-1/

# SAMPLE SET 2
git clone git@github.com:NYULibraries/dlts-finding-aids-ead-sample-set-2.git

find dlts-finding-aids-ead-sample-set-2/ -type f -name *.xml -exec cat {} \; | zek | grep -v '// Ead was generated' > sample-set-2-ead-files.go

rm sample-set-2-individual-ead-files/*

for eadfile in $( find dlts-finding-aids-ead-sample-set-2/ -type f -name *.xml ); do eadid=$( basename $eadfile .xml ); cat $eadfile | zek | grep -v '// Ead was generated' > sample-set-2-individual-ead-files/${eadid}.go; done

rm -fr dlts-finding-aids-ead-sample-set-2/

