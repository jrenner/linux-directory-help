#!/bin/bash
go-linux-amd64 build -o /home/jrenner/projects/dirhelp/bin/dirhelp-linux-amd64 dirhelp.go
go-linux-386 build -o /home/jrenner/projects/dirhelp/bin/dirhelp-linux-386 dirhelp.go
go-linux-arm build -o /home/jrenner/projects/dirhelp/bin/dirhelp-linux-arm dirhelp.go