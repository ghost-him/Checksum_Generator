# Checksum_Generator
**用于计算文件md5, sha1, sha256, sha512的命令行程序**



---

**用法:**

```
.\Checksum_Generator.exe
```

```
-f string
        文件的路径 绝对路径或者是相对路径
-mode string
        加密的方式 md5 / sha1 / sha256 / sha512 / all (default "md5")
-o string
        结果的输出路径, 若不填则不生成文件
```



**示例**

1. 计算同目录下 “input.txt” 文件的 md5值, 结果不输出到文件

```
.\Checksum_Generator.exe -f input.txt
```

2. 计算同目录下 “input.txt” 文件的 sha1值并输出到“answer.txt”文件中

```
.\Checksum_Generator.exe -f input.txt -mode sha1 -o answer.txt
```

3. 计算同目录下 “game.txt” 文件所有的值并输出到“answer.txt”文件中

```
.\Checksum_Generator.exe -f game.txt -mode all -o answer.txt
```

4. 计算桌面上的“test.txt”文件的sha256的值, 结果不输出到文件

```
.\Checksum_Generator.exe -f C:\Users\ghost\Desktop\test.txt -mode sha256
```

5. 计算同目录下的“test.txt”文件的sha512的值, 不输出文件

```
.\Checksum_Generator.exe -f test.txt -mode sha512
```

## 更新日记

2022/3/16

* 完成程序的基本的功能

2022/3/20

* 添加路径校验,可以直接双击打开程序了,不需要通过命令行参数来指定路径

* 允许在窗口显示计算结果,不另外创建文件
