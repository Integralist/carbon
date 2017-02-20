FROM golang:1.7.5-alpine

RUN apk --update add vim git

COPY .vim /root/.vim
COPY .vimrc /root/.vimrc

RUN vim -c "GoInstallBinaries"
