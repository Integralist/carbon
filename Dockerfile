FROM golang:1.7.5-alpine

RUN apk --update add vim git

COPY .vim /root/.vim
COPY .vimrc /root/.vimrc

# Use vim's execute command to pipe commands so we avoid "Press ENTER or type command to continue"
RUN vim -c "execute 'silent GoUpdateBinaries' | execute 'quit'"
