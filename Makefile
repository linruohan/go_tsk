all:win

win:
	go build -tags tempdll -ldflags "-s -w -H windowsgui" && ./*.exe