build:
	docker build -t go-container-with-vim .

run: build
	docker run -it -v "$$HOME/.vimrc":/root/.vimrc \
		             -v "$$HOME/.vim":/root/.vim \
								 -v "$$(pwd)":/go/src go-container-with-vim /bin/ash

gobuild:
	go build ./carbon.go

install: gobuild
	cp ./carbon  /usr/local/bin/carbon
	rm ./carbon

clean:
	docker rmi go-container-with-vim &> /dev/null
	rm ./carbon &> /dev/null
