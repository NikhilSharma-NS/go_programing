package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufpool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocatuion of bytes buffer")
		return new(bytes.Buffer)
	},
}

func main() {

	logs(os.Stdout, "debug-st")
	time.Sleep(10 * time.Microsecond)
	logs(os.Stdout, "debug-st1")
	time.Sleep(10 * time.Microsecond)
	logs(os.Stdout, "debug-st2")

	time.Sleep(10 * time.Microsecond)
	logs(os.Stdout, "debug-st3")

	time.Sleep(10 * time.Microsecond)
	logs(os.Stdout, "debug-st4")

	time.Sleep(10 * time.Microsecond)
	logs(os.Stdout, "debug-st5")

}

// create pool of bytes.Buffers which can be used

func logs(w io.Writer, debug string) {
	//var b bytes.Buffer
	b := bufpool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(debug)
	b.WriteString("\n")
	w.Write(b.Bytes())
	bufpool.Put(b)
}
