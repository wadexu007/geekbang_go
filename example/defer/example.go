package main

import (
	"io"
	"os"
)

func main() {

}

//这段代码可以运行，但存在'安全隐患'。如果调用dst, err := os.Create(dstName)失败，
//则函数会执行return退出运行。但之前创建的src(文件句柄)没有被释放。
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

//使用defer则可以避免这种情况的发生
func CopyFileDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
