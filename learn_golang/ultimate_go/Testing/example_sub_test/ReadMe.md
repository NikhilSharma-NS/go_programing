go test -v
go test -run TestSub
go test -run TestSub/ok -v
go test -run TestSub/notfound -v


go test -run TestSub -v
go test -run TestSub/statusok -v
go test -run TestSub/statusnotfound -v
go test -run TestParallelize -v

