package main

import (
	"flag" // command line option parser
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

// 命令行 选项参数 获取
// go run command_flag.go A B C 输出 "A B C"
// go run command_flag.go -n A B C 输出 "A 换行 B 换行 C 换行"
func main() {
	// 默认打印选项提示
	flag.PrintDefaults()

	flag.Parse() // Scans the arg list and sets up flags

	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}

	os.Stdout.WriteString(s)
}