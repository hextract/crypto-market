#!/bin/bash

cd ../

cd fetcher
rm  internal/di/wire_gen.go
wire gen internal/di/*.go
cd ..

cd receiver
rm  internal/di/wire_gen.go
wire gen internal/di/*.go
cd ..

cd reporter
rm  internal/di/wire_gen.go
wire gen internal/di/*.go
cd ..

cd executor
rm  internal/di/wire_gen.go
wire gen internal/di/*.go
cd ..

cd generator
rm  internal/di/wire_gen.go
wire gen internal/di/*.go
cd ..
