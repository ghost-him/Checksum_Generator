package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var file_path string
var mode string
var file_output_path string
var now_time string

var abs_path string
var file_name string

var (
	md5_code    string
	sha1_code   string
	sha256_code string
	sha512_code string
)

func init() {
	flag.StringVar(&file_path, "f", "", "文件的路径 绝对路径或者是相对路径")
	flag.StringVar(&mode, "mode", "md5", "加密的方式 md5 / sha1 / sha256 / sha512 / all")
	flag.StringVar(&file_output_path, "output_path", "output.txt", "结果的输出路径")
	flag.Parse() // 读取命令行
	// 获取当前文件的目录
	if filepath.IsAbs(file_path) {
		abs_path = filepath.Dir(file_path)
		file_name = file_path[strings.Index(file_path, abs_path)+len(abs_path)+1:]
	} else {
		file_name = file_path
		abs_path, _ = filepath.Abs(file_path)
		abs_path = filepath.Dir(abs_path)
	}
	abs_path += "\\"
	// 获取时间
	now_time = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}

// 获取文件的md5码
func file_md5(file_path *os.File) string {
	file_path.Seek(0, 0)
	md5h := md5.New()
	io.Copy(md5h, file_path)
	return hex.EncodeToString(md5h.Sum(nil))
}

// 获取文件的sha1码
func file_sha1(file_path *os.File) string {
	file_path.Seek(0, 0)
	sha1h := sha1.New()
	io.Copy(sha1h, file_path)
	return hex.EncodeToString(sha1h.Sum(nil))
}

// 获取文件的sha256码
func file_sha256(file_path *os.File) string {
	file_path.Seek(0, 0)
	sha256h := sha256.New()
	io.Copy(sha256h, file_path)
	return hex.EncodeToString(sha256h.Sum(nil))
}

// 获取文件的sha512码
func file_sha512(file_path *os.File) string {
	file_path.Seek(0, 0)
	sha512h := sha512.New()
	io.Copy(sha512h, file_path)
	return hex.EncodeToString(sha512h.Sum(nil))
}

func file_write(file_path_ *os.File) {
	file_path_.WriteString("当前的时间: " + now_time + "\n")
	file_path_.WriteString("注意：如果字符串中包含有回车(\\r)换行符(\\n)等，因为不同的系统存在差异则会导致加密结果不同。所以最好将字符串以你需要的格式保存为文件后再选择文件加密！\n")
	output_text := ""
	output_text += "文件路径: " + abs_path + file_name + "\n"
	output_text += "文件名字: " + file_name
	output_text += "\n\n"

	if md5_code != "" {
		output_text += "md5(32位小写): " + md5_code + "\n"
		output_text += "md5(32位大写): " + strings.ToUpper(md5_code)
	}
	output_text += "\n"

	if sha1_code != "" {
		output_text += "sha1: " + sha1_code
	}
	output_text += "\n"

	if sha256_code != "" {
		output_text += "sha256: " + sha256_code
	}
	output_text += "\n"

	if sha512_code != "" {
		output_text += "sha512: " + sha512_code
	}
	output_text += "\n"
	file_path_.WriteString(output_text)
}

// 计算
func compute(input_file_path, output_file_path *os.File) {
	switch {
	case mode == "md5":
		md5_code = file_md5(input_file_path)
	case mode == "sha1":
		sha1_code = file_sha1(input_file_path)
	case mode == "sha256":
		sha256_code = file_sha256(input_file_path)
	case mode == "sha512":
		sha512_code = file_sha512(input_file_path)
	case mode == "all":
		{
			md5_code = file_md5(input_file_path)
			sha1_code = file_sha1(input_file_path)
			sha256_code = file_sha256(input_file_path)
			sha512_code = file_sha512(input_file_path)
		}
	}

	file_write(output_file_path)
}

func main() {

	// 打开输出的文件
	output_file, _ := os.OpenFile(file_output_path, os.O_CREATE|os.O_RDWR, 0666) // 创建输出文件
	defer output_file.Close()                                                    // 关闭文件

	// 检测输入的文件
	input_path, err := os.Open(file_path)
	if err != nil {
		output_text := fmt.Sprintf("%s \n输入文件读取失败, 当前的文件路径为 %s", now_time, abs_path+file_name)
		output_file.WriteString(output_text)
		return
	}
	compute(input_path, output_file)
}
