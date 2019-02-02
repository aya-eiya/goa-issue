#!/bin/bash

rm -rf cmd/ gen/ 
goa gen github.com/aya-eiya/goa-issue/api/design
goa example github.com/aya-eiya/goa-issue/api/design
cd cmd/calc/
go build
./calc
