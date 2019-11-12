package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/gosuri/uilive"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	og := []string{
		"   D;;:;;:;;:;;:;:;:;;:;:;:;:;;:;;:;;:;;:;z  ",
		"   $;:;:;::;::;:;;:;;:;;:;;:;;::;;::;;:;;;I  ",
		"  .I;;:;;:;;:;;::;;:;:;:;;:;:;;:;:;:;::;;:I  ",
		"  ,<;;::;:;;::;;:;;;;;;;;:;::;;:;;:;;;:;;;I  ",
		"  ,(;;;:;::;:;;::;;j=1J71<;;;:;:;;::;:;:;:I  ",
		"  J;;:;;;:;;::;;;;:r  ] .>;;;:;:;:;;:;:;;;r  ",
		"  z;;::;:;;:;;:;;j=<?75?7~?I;;:;;:;;;:;:;<]  ",
		"  (<;;;;;;:;;;;;;?+~(J-J-_(3;;;;;;::;;:;;+|  ",
		"  ,(;:;:;j/7!''??1+?MMMMM1+?7771+<;;;:;;:j   ",
		"  .P;;;;J!..       4;<<iJ        .4<;;:;;2   ",
		".3;J<;;j|(M#Q       D;<2.MM5.      1:;;;j73, ",
		"$;jMN<;?|,WH3       $;:t.MM#       ,(;;jP;;?|",
		"4<;T9TJ;?.        .J;;;?&         .t;;jM@:;+%",
		" (1++++Y+;?C+...J7<;;;:;;?i..  ..J>;jv<;;;j= ",
		"         .71+<;;;;;;;:;;;;;;;;;;<+J=  ?77!   ",
		"             '_?771+++++++++?77!             ",
	}
	ogw := len(og[0])

	tw, _, err := terminal.GetSize(syscall.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i := range og {
		og[i] = strings.Repeat(" ", tw) + og[i]
	}

	writer := uilive.New()

	writer.Start()
	for i := 0; i <= tw+ogw; i++ {
		var out string
		for _, v := range og {
			if i < ogw {
				out += v[i:tw+i] + "\n"
			} else if i < len(v) {
				out += v[i:] + "\n"
			} else {
				out += "\n"
			}
		}
		fmt.Fprint(writer, out)
		time.Sleep(10 * time.Millisecond)
	}
	writer.Stop()
}
