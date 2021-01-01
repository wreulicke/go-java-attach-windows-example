FROM golang:1.15

WORKDIR /opt/app

RUN apt update && apt install -y --no-install-recommends gcc-mingw-w64 gcc-mingw-w64-x86-64 pkg-config \
  && rm -rf /var/lib/apt/lists/*
ADD . .
RUN x86_64-w64-mingw32-gcc -c -DJATTACH_VERSION=\\"1.5\\" -o jattach.o jattach.c && x86_64-w64-mingw32-gcc -shared -o jattach.dll jattach.o
RUN GO111MODULES=on GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build .
