#### wait for Result

The wait for result pattern is a foundational pattern used by larger patterns like fan
out/in. In this pattern, a Goroutine is created to perform some known work and
signals their result back to the Goroutine that created them. This allows for the
actual work to be placed on a Goroutine that can be terminated or walked away
from. 