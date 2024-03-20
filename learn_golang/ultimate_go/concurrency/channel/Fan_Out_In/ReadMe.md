The fan out/in pattern uses the wait for result pattern just described.


I must remember, a fan out is dangerous in a running service since the number of
child Goroutines I create for the fan are a multiplier. If I have a service handling
50k requests on 50 thousand Goroutines, and I decide to use a fan out pattern of
10 child Goroutines for some of the requests, in a worse case scenario I would be
talking 500k Goroutines existing at the same time. Depending on the resources
those child Goroutines needed, I might not have them available at that scale and
the back pressure could bring the service down.
