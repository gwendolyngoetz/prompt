build: clean
ifeq ($(OS),Windows_NT)
	@go build -ldflags="-X 'main.Version=v$(VERSION)'" -o .\build\promptwin.exe .\cmd\prompt\main.go
else
	@go build -ldflags="-X 'main.Version=v$(VERSION)'" -o build/prompt cmd/prompt/main.go
	@GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.Version=v$(VERSION)'" -o ./build/promptwin.exe cmd/prompt/main.go
endif
	@echo Built prompt

clean:
ifeq ($(OS),Windows_NT)
	@if exist ".\build" rmdir .\build /s /q 
else
	@rm -rf build
endif
	@echo Cleaned

install: uninstall
ifeq ($(OS),Windows_NT)
	@if not exist "%LOCALAPPDATA%\prompt" mkdir %LOCALAPPDATA%\prompt
	@copy .\build\promptwin.exe %LOCALAPPDATA%\prompt /y
else
	@[ -d $${HOME}/.local/bin ] || @mkdir -p $${HOME}/.local/bin
	@cp build/prompt $${HOME}/.local/bin
endif
	@echo Installed prompt

uninstall:
ifeq ($(OS),Windows_NT)
	@if exist "%LOCALAPPDATA%\prompt" rmdir "%LOCALAPPDATA%\prompt" /s /q 
else
	@rm $${HOME}/.local/bin/prompt
endif
	@echo Uninstalled prompt

define newline


endef

define CONTROL_FILE_BODY
Package: prompt
Version: $(VERSION)
Section: base
Priority: optional
Architecture: x64
Maintainer: Gwendolyn Goetz
Description: Fancy prompt output
endef

package-deb:
	@mkdir -p ./package/prompt_$(VERSION)/tmp/usr/local/bin
	@mkdir -p ./package/prompt_$(VERSION)/DEBIAN

	@cp ./build/prompt ./package/prompt_$(VERSION)/tmp/usr/local/bin
	@touch ./package/prompt_$(VERSION)/DEBIAN/control
	
	@echo '$(subst $(newline),\n,${CONTROL_FILE_BODY})' > ./package/prompt_$(VERSION)/DEBIAN/control

	@dpkg-deb --build ./package/prompt_$(VERSION)
	@echo Package deb
