APP_NAME := genpass

build:
	echo "Compiling for every OS and Platform"
	$(foreach OS, windows linux darwin, \
		$(foreach ARCH, amd64 arm64, \
			$(if $(filter $(OS),windows), $(eval EXT=.exe), $(eval EXT=)) \
				GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags="-s -w" -o bin/$(APP_NAME)_$(OS)_$(ARCH)$(EXT) main.go; \
			))
