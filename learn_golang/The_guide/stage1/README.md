### Simple Start 

#### stage 1 

1) hello world program 
2) Basic Question
   2.1) How do we run the code in our project
   2.2) What does package mean
   2.3) what does import fmt mean 
   2.4) what that func thing
   2.5) how is the main.go file organized

2.1) 
```
go build -> compile the file
go run -> compile and execute
go fmt -> format all the code in the current directory
go install -> compile and install a package
go get -> downloads the raw source code of someone else's package
go test -> run the any tests associated with current project
```

2.2) 
Package == Project == workspace

```
main.go                       support.go
package main              package main
```
Types of packages

1) Executable -> generate a file that we can run
2) Reusable -> Code used as 'helpers' good place to put reusable logic

 ```
package main -> defines a packages that can be compiled and then executed (must have a func called main)

package cal
                    -> defines a package that can be used as a dependency (helper code)
package upload 
 ```

2.3) what does import fmt mean 

standard libary of go 
like : debug,io,crypto,encoding,math

import is to use standarad and dependent package import


2.4) what that func thing

```
func                                     main                             ()

tells go we're                        set the name of func            list of argument to pass the function 
about to declare a func

{
    // function body
}
```

2.5) how is the main.go file organized

```
package main -> package declaration 
import "fmt" -> import other packages
func main(){
             -> declare functions , tell Go to do things
}
```

