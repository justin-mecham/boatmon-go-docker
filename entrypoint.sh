#!/bin/sh

go-wrapper download
go-wrapper install
exec go-wrapper run  #exec is important for having NOT BASH running as primary PID
