

target = toy_docker

all: $(target)
	echo "build $(target)"

$(target):
	go build -o toy_docker

2.2.3:
	sudo go run ch2.3.3.go

clean:
	rm -f $(target)

