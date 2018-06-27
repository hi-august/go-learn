package main

import (
	"fmt"
	"os"
	// 带缓冲的io处理
	"bufio"
	// 底层io功能
	"io"
	// log日志相关
	"log"
	"path/filepath"
)

func main() {
	infileName, outfileName, err := loadFromCmd()
	// fmt.Println(infileName, outfileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	infile, outfile := os.Stdin, os.Stdout
	if infileName != "" {
		if infile, err = os.Open(infileName); err != nil {
			log.Fatal(err)
		}
		defer infile.Close()
	}
	if outfileName != "" {
		if oufile, err = os.Create(outfileName); err != nil {
			log.Fatal(err)
		}
		defer outfile.Close()
	}
	if err = process(infile, outfile); err != nil {
		log.Fatal(err)
	}
}

// 从命令行读取参数
func loadFromCmd() (infileName, outfileName string, err error) {
	// 提示信息
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<] infile.txt [>] outfile.txt", filepath.Base(os.Args[0]))
		return "", "", err
	}
	if len(os.Args) > 1 {
		outfileName = os.Args[2]
	}
	// in和out同一个处理
	if infileName != "" && infileName == outfileName {
		log.Fatal("won't overwrite the infile")
	}
	return infileName, outfileName, nil
}

// 处理具体
func process(infil io.Reader, outfile io.Writer) (err error) {
	reader := bufio.NewReader(infile)
	writer := bufio.NewWriter(oufile)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}
