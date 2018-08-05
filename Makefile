

target = toy_docker

all: $(target)
	echo "build $(target)"

$(target):
	go build -o toy_docker


clean:
	rm -f $(target)

