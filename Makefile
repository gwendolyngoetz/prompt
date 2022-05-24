build: clean
ifeq ($(OS),Windows_NT)
	@go build -o .\build\promptwin.exe .\cmd\prompt\main.go
else
	@go build -o build/prompt cmd/prompt/main.go
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




 