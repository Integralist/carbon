.DEFAULT_GOAL := install

copy_vim_files:
	cp -r "$$HOME/.vim" ./.vim
	cp "$$HOME/.vimrc" ./.vimrc

remove_vim_files:
	-@rm -rf ./.vim &> /dev/null || true
	-@rm -rf ./.vimrc &> /dev/null || true

build: copy_vim_files
	docker build -t go-container-with-vim .
	remove_vim_files

run: build
	docker run -it \
		-v "$$HOME/.vimrc":/root/.vimrc \
		-v "$$HOME/.vim":/root/.vim \
		-v "$$(pwd)":/go/src go-container-with-vim /bin/ash

gobuild:
	go build ./carbon.go

install: gobuild
	cp ./carbon  /usr/local/bin/carbon
	rm ./carbon

# -       allows errors to be ignored (i.e. don't stop further execution steps)
# @       stops makefile from printing command that was executed
# &>      redirects stdout/stderr to /dev/null (as command can sometimes error)
# || true prevents makefile from printing 'error ignored'
#
# I had hoped .SILENT: clean would've sufficed but alas it was not to be

clean:
	-@docker rmi -f go-container-with-vim &> /dev/null || true
	-@rm ./carbon &> /dev/null || true

uninstall: clean
	-@rm /usr/local/bin/carbon &> /dev/null || true
