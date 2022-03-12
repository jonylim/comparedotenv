#!/bin/bash

if [ -z $1 ]; then
	CGO_ENABLED=0 go build .
else
	case  "$1" in
		"linux" )
			CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
			;;
		* )
			echo arg is invalid
			;;
	esac
fi
