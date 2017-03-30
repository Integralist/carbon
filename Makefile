# -       allows errors to be ignored (i.e. don't stop further execution steps)
# @       stops makefile from printing command that was executed
# &>      redirects stdout/stderr to /dev/null (as command can sometimes error)
# || true prevents makefile from printing 'error ignored'

copy_vim_files:
	@if [ ! -d "./.vim" ]; then cp -r "$$HOME/.vim" ./.vim; fi
	@if [ ! -f "./.vimrc" ]; then cp "$$HOME/.vimrc" ./.vimrc; fi

remove_vim_files:
	@rm -rf ./.vim
	@rm ./.vimrc

build: copy_vim_files
	@docker build -t go-container-with-vim .

run: build
	@docker run -it -v "$$(pwd)":/go/src go-container-with-vim /bin/ash

clean: remove_vim_files
	-@docker rmi -f go-container-with-vim &> /dev/null || true
	-@rm ./app &> /dev/null || true

rebuild: clean run

uninstall: clean
	-@rm /usr/local/bin/app &> /dev/null || true

gobuild:
	go build -o ./carbon

install: gobuild
	cp ./carbon  /usr/local/bin/carbon
	rm ./carbon

compile:
	@docker build -t go-compiler -f ./Dockerfile-compile .
	@docker run -it -v "$$(pwd)":/go/src go-compiler || true
