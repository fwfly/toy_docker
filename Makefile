
.PHONY: ch3 clean

target = toy_docker

all: $(target)
	echo "build $(target)"

$(target):
	go build -o toy_docker

2.2.3:
	sudo go run ch2.3.3.go

ch3_module:
	cp -r ch3/container  ~/go/src/ch3/

ch3: ch3_module
	sudo go build ch3

clean:
	rm -rf ~/go/src/ch3
	rm -f ch3/ch3
	rm -f $(target)

