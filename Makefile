build: clean
	@VERSION=$(VERSION) ./bin/build.sh
	echo Built prompt

clean:
	@rm -rf ./build

install: uninstall
	@[ -d $${HOME}/.local/bin ] || @mkdir -p $${HOME}/.local/bin
	@cp build/linux/amd64/prompt $${HOME}/.local/bin
	@echo Installed prompt

uninstall:
	@rm -f $${HOME}/.local/bin/prompt
	@echo Uninstalled prompt

package-deb:
	@VERSION=$(VERSION) ./bin/package-deb.sh
	@echo Packaged Deb

package-zip:
	@VERSION=$(VERSION) ./bin/package-win.sh
	@echo Packaged Windows

package: package-deb package-zip
