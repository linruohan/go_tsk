#!/usr/bin/env bash
# x64
windres.exe -i app.rc -o app.syso -F pe-x86-64
# x86
#windres.exe -i app.rc -o app.syso -F pe-i386
# -s：忽略符号表和调试信息。
# -w：忽略DWARFv3调试信息，使用该选项后将无法使用gdb进行调试。
go build -tags tempdll -ldflags "-s -w -H windowsgui" -o goTsk.exe
upx -9 ./*.exe
./*.exe