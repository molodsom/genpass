APP_NAME := genpass

build:
	echo "Compiling for every OS and Platform"
	$(foreach GOOS, windows linux darwin, \
		$(foreach GOARCH, 386 amd64 arm64, \
			$(if $(filter $(GOOS),windows), $(eval EXT=.exe), $(eval EXT=)) \
				GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(APP_NAME)_$(GOOS)_$(GOARCH)$(EXT) main.go;))
