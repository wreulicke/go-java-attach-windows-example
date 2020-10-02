
.PHONY: all clean

all:
	gcc -c -DJATTACH_VERSION=\"1.5\" -o jattach.o jattach.c
# On windows platforms it's a `.dll` and there's no leading `lib`
	gcc -shared -o jattach.dll jattach.o
	go build main.go

clean:
	rm jattach.o jattach.dll