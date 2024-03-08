CPU Profiling

$ go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out

$ go tool pprof p.out

(pprof) list algOne
(pprof) web list algOne


$ go tool pprof -http :3000 p.out


Memory Profiling

$ go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out

go tool pprof p.out

(pprof) list algOne
(pprof) web list algOne