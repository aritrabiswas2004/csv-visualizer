build:main.go
	go build -o viz main.go

test:
	go test -v

clean:
	rm -rf viz

# This is just like go run main.go
# I am only keeping it temporary (remove in stable because of hardcoded .csv path)
# Possibly in this file only for the vibes
# Avoids generating executable (creates in /tmp)
run:main.go
	go run main.go
