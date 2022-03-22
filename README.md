# Checksum_Generator
**用于计算文件md5, sha1, sha256, sha512的命令行程序**



---

**用法:**

```
.\Checksum_Generator.exe
```

```
  -d string
        文件夹的路径,会循环遍历所有的文件, 如果为空,则读取文件的路径
  -f string
        文件的路径 绝对路径或者是相对路径
  -m string
        加密的方式 md5 / sha1 / sha256 / sha512 / all 通过 , 来实现多种模式的组合 (default "md5")
  -o string
        结果的输出路径, 若不填则不生成文件
```



**示例**

1. 计算同目录下 “input.txt” 文件的 md5值, 结果不输出到文件

```
.\Checksum_Generator.exe -f input.txt
```

2. 计算同目录下 “input.txt” 文件的 sha1值和sha512值并输出到“answer.txt”文件中

```
.\Checksum_Generator.exe -f input.txt -m sha1,sha512 -o answer.txt
```

3. 计算同目录下 “game.txt” 文件所有的值并输出到“answer.txt”文件中

```
.\Checksum_Generator.exe -f game.txt -m all -o answer.txt
```

4. 计算桌面上的“test.txt”文件的sha256的值, 结果不输出到文件

```
.\Checksum_Generator.exe -f C:\Users\ghost\Desktop\test.txt -m sha256
```

5. 计算同目录下的“test.txt”文件的sha512的值, 不输出文件

```
.\Checksum_Generator.exe -f test.txt -m sha512
```

6. 计算桌面上所有文件的sha1值,结果不输出到文件中

```
.\Checksum_Generator.exe -d C:\Users\ghost\Desktop -m sha1
```



## 更新日记

2022/3/16

* 完成程序的基本的功能

2022/3/20

* 添加路径校验,可以直接双击打开程序了,不需要通过命令行参数来指定路径

* 允许在窗口显示计算结果,不另外创建文件

2022/3/22

* 添加文件夹模式,可以直接计算指定文件夹下的所有文件的加密值
