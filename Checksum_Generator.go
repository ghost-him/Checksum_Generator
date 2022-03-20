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

var file_path string                             // 文件的路径
var mode map[string]bool = make(map[string]bool) // 文件的模式
var file_output_path string                      // 输出文件的路径
var now_time string                              // 现在的时间

var is_file_path_true bool = false // 文件的路径是否正确
var is_spawn_file bool = false     // 是否生成文件

var abs_path string  // 文件的绝对路径
var file_name string // 文件的名字
var mode_text string

var (
	md5_code    string
	sha1_code   string
	sha256_code string
	sha512_code string
)

func init() {
	var slice = []string{} // 临时的切片
	var modeT string       // 临时的变量
	flag.StringVar(&file_path, "f", "", "文件的路径 绝对路径或者是相对路径")
	flag.StringVar(&modeT, "mode", "md5", "加密的方式 md5 / sha1 / sha256 / sha512 / all 通过 , 来实现多种模式的组合")
	flag.StringVar(&file_output_path, "o", "", "结果的输出路径, 若不填则不生成文件")
	flag.Parse() // 读取命令行

	slice = strings.Split(modeT, ",")
	for i := 0; i < len(slice); i++ {
		mode[slice[i]] = true
	}

	// 获取时间
	now_time = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	if file_output_path == "" {
		is_spawn_file = false
	} else {
		is_spawn_file = true
	}
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

func answer_write(file_path_ *os.File) {
	output_text := "当前的时间: " + now_time + "\n"
	output_text += "注意：如果字符串中包含有回车(\\r)换行符(\\n)等，因为不同的系统存在差异则会导致加密结果不同。所以最好将字符串以你需要的格式保存为文件后再选择文件加密！\n"
	output_text += "文件路径: " + abs_path + file_name + "\n"
	output_text += "文件名字: " + file_name
	output_text += "当前的模式: " + mode_text
	output_text += "\n\n"

	if md5_code != "" {
		output_text += "md5(32位小写): " + md5_code + "\n"
		output_text += "md5(32位大写): " + strings.ToUpper(md5_code) + "\n"
	}

	if sha1_code != "" {
		output_text += "sha1: " + sha1_code + "\n"
	}

	if sha256_code != "" {
		output_text += "sha256: " + sha256_code + "\n"
	}

	if sha512_code != "" {
		output_text += "sha512: " + sha512_code + "\n"
	}

	if is_spawn_file {
		file_path_.WriteString(output_text)
	} else {
		fmt.Println(output_text)
		var temp int
		fmt.Scan(&temp)
	}
}

// 计算
func compute(input_file_path, output_file_path *os.File) {
	if mode["all"] {
		mode_text += "md5 sha1 sha256 sha512"
		md5_code = file_md5(input_file_path)
		sha1_code = file_sha1(input_file_path)
		sha256_code = file_sha256(input_file_path)
		sha512_code = file_sha512(input_file_path)
	} else {
		if mode["md5"] {
			mode_text += "md5 "
			md5_code = file_md5(input_file_path)
		}
		if mode["sha1"] {
			mode_text += "sha1 "
			sha1_code = file_sha1(input_file_path)
		}
		if mode["sha256"] {
			mode_text += "sha256 "
			sha256_code = file_sha256(input_file_path)
		}
		if mode["sha512"] {
			mode_text += "sha512 "
			sha512_code = file_sha512(input_file_path)
		}
	}
	answer_write(output_file_path)
}

func get_file_name() {

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
}

func check_file(file_path string) *os.File {
	input_fileP, err := os.Open(file_path)
	if err != nil {
		is_file_path_true = false
	} else {
		is_file_path_true = true
	}
	return input_fileP
}

func main() {
	var output_fileP *os.File // 输出文件的指针
	var input_fileP *os.File  // 输入文件的指针
	defer output_fileP.Close()
	defer input_fileP.Close()

	// 检测输入的文件
	input_fileP = check_file(file_path)
	get_file_name()
	for !is_file_path_true {
		file_path = ""
		fmt.Println("检测不到文件,当前的文件路径是 " + abs_path + file_name + "\n请输入文件路径")
		fmt.Scanln(&file_path)
		input_fileP = check_file(file_path)
		get_file_name()
	}

	// 打开输出的文件
	if is_spawn_file {
		output_fileP, _ = os.OpenFile(file_output_path, os.O_CREATE|os.O_RDWR, 0666) // 创建输出文件
		defer output_fileP.Close()                                                   // 关闭文件
	}

	compute(input_fileP, output_fileP)

}
