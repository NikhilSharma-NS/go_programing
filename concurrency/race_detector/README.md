#### Race Detector

a) go provides race detector tool for finding race conditions in Go code

- go test -race mypkg // test the package
- go run -race mysrc.go // complte and run the program
- go build -race mycmd // build the command
- go install -race mypkg // install the package

b) binary needs to be race enabled

c) when racy behavior is detected a warning is printed

d) Race enabled binary will 10 times slower and consume 10 times more memory

e) Integration tests and load tests are good candidate to test with binary with race enabled

