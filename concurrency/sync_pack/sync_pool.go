package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

func main11() {

	log(os.Stdout, "debug-st")
	time.Sleep(10 * time.Microsecond)
	log(os.Stdout, "debug-st1")
	time.Sleep(10 * time.Microsecond)
	log(os.Stdout, "debug-st2")

	time.Sleep(10 * time.Microsecond)
	log(os.Stdout, "debug-st3")

	time.Sleep(10 * time.Microsecond)
	log(os.Stdout, "debug-st4")

	time.Sleep(10 * time.Microsecond)
	log(os.Stdout, "debug-st5")

}

func log(w io.Writer, debug string) {
	var b bytes.Buffer
	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(debug)
	b.WriteString("\n")
	w.Write(b.Bytes())
}
