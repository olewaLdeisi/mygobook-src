package main

import (
	"fmt"
	"io"
	"os"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

// 告诉编译器PathError可以作为error传递
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()
	/*
		defer func() {
			// 多句
		} ()
	*/
	return io.Copy(dstFile, srcFile)
}

func t(a int) {
	if a < 5 {
		return
	}
	defer func() {
		fmt.Println("---------------")
	}()
	fmt.Println("normal ", a)
}

func main() {
	fmt.Println("==============")
	t(5)
	fmt.Println("==============")
	t(4)
	fmt.Println("==============")
	t(9)
}
