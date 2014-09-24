


all: fmt
	go build .
	. environment

fmt:
	go fmt *.go

clean:
	rm -f halbot
	@echo unset HAL_HIPCHAT_USER
	@echo unset HAL_HIPCHAT_PASSWORD
	@echo unset HAL_HIPCHAT_ROOMS
	@echo unset HAL_HIPCHAT_RESOURCE
	@echo unset HAL_ADAPTER
	@echo unset HAL_LOG_LEVEL
