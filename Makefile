


all: fmt
	go build .
	. environment

fmt:
	go fmt *.go

clean:
	rm -f mybot
	unset HAL_HIPCHAT_USER
	unset HAL_HIPCHAT_PASSWORD
	unset HAL_HIPCHAT_ROOMS
	unset HAL_HIPCHAT_RESOURCE
	unset HAL_ADAPTER
	unset HAL_LOG_LEVEL
