The pooling pattern uses the wait for task pattern just described. The pooling
pattern allows me to manage resource usage across a well defined number of
Goroutines. 

As explained previously, in Go pooling is not needed for efficiency in
CPU processing like at the operating system. It’s more important for efficiency in
resource usage.
