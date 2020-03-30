# Big Data
## Meeting with Hadoop
### 2020.3.28 [log]
#### Activity
- VMware 搭建虚拟机，虚拟环境配置
- Hadoop 安装，Example测试（standalone模式）
#### Traps
- 执行文件目录和安装文件目录的区别
	- ls -lrt 寻找文件之间的连接关系
	- chown chmod
- JAVA_HOME配置
- 软件教程一定要注意其版本
- ./output vs output 本地文件系统与HDFS的区别
### 2020.3.29 [log]
#### Activity
- Hadoop 伪分布式
- ssh远程登陆 scp远程传输
#### Traps
- sshd.config注意是否禁用了密码登陆
### 2020.3.28 [log]
#### Activity

#### Traps
- go module 的使用（1.11～ *， 1.14 默认开启GO111MODULE=on）
- go env -w 不能与环境变量冲突
- [golang]	strings.Split() by '\t' 的切分迷之看不懂，迷之出错
- [golang]	处理字符串时一定要注意编码的不同
``` golang
func scanLine() (string, error) {
	var words []rune	//采用rune符合unicode编码
	var word rune
	var err error = nil
	for {
		_, err = fmt.Scanf("%c", &word)
		if err != nil || word == '\n' {
			break
		}
		words = append(words, word)
	}
	if err == io.EOF && len(words) > 0 {
		err = nil	//防止没有换行直接抛出EOF
	}
	return string(words), err
}
```
