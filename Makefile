build:
	@go build -o build/prompt cmd/prompt/main.go
	@echo Built prompt

clean:
	@rm -rf build
	@echo Cleaned build directory
	
install:
	@[ -d $${HOME}/.local/bin ] || @mkdir -p $${HOME}/.local/bin
	@cp build/prompt $${HOME}/.local/bin
	@echo Installed prompt

uninstall:
	@rm $${HOME}/.local/bin/prompt
	@echo Uninstalled prompt
