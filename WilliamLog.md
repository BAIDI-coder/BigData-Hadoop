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
### 2020.4.1 [log]
#### Activity
- hadoop集群搭建（两台虚拟机，一台运行NameNode，一台运行DataNode）
``` shell
[hadoop@localhost hadoop]$ hdfs dfsadmin -report
2020-04-01 23:26:08,129 WARN util.NativeCodeLoader: Unable to load native-hadoop library for your platform... using builtin-java classes where applicable
Configured Capacity: 18238930944 (16.99 GB)
Present Capacity: 11914702848 (11.10 GB)
DFS Remaining: 11914698752 (11.10 GB)
DFS Used: 4096 (4 KB)
DFS Used%: 0.00%
Replicated Blocks:
	Under replicated blocks: 0
	Blocks with corrupt replicas: 0
	Missing blocks: 0
	Missing blocks (with replication factor 1): 0
	Low redundancy blocks with highest priority to recover: 0
	Pending deletion blocks: 0
Erasure Coded Block Groups: 
	Low redundancy block groups: 0
	Block groups with corrupt internal blocks: 0
	Missing block groups: 0
	Low redundancy blocks with highest priority to recover: 0
	Pending deletion blocks: 0

-------------------------------------------------
Live datanodes (1):

Name: 192.168.100.11:9866 (hadoop-worker.com)
Hostname: hadoop-worker.com
Decommission Status : Normal
Configured Capacity: 18238930944 (16.99 GB)
DFS Used: 4096 (4 KB)
Non DFS Used: 6324228096 (5.89 GB)
DFS Remaining: 11914698752 (11.10 GB)
DFS Used%: 0.00%
DFS Remaining%: 65.33%
Configured Cache Capacity: 0 (0 B)
Cache Used: 0 (0 B)
Cache Remaining: 0 (0 B)
Cache Used%: 100.00%
Cache Remaining%: 0.00%
Xceivers: 1
Last contact: Wed Apr 01 23:26:09 CST 2020
Last Block Report: Wed Apr 01 23:20:30 CST 2020
Num of Blocks: 0
```

#### Traps
- ssh对于.ssh内的文件有权限限制（不允许组有写权限等）
- ssh登陆用户如果不指定则默认同名用户，因此要注意用户名问题（实验中将两台虚拟机运行hadoop的用户名设为相同也是处于这种考虑）

### 2020.4.2 [log]
### Activity
- 用hadoop执行自己的mapred streaming job

### Traps
- 执行mapred streaming时显示connection refused
	- solution: 虚拟网卡异常，重启恢复
- 执行mapred streaming时抛出找不到class异常
	- solution: 根据诊断信息发现mapred-site.xml文件没有配置相关value
- map succeed but reduce failed
	- solution: 将reducer暂时改为/usr/bin/cat，发现reduce input的key/value pair分隔符与脚本程序默认分隔符不一致
	- solution: 添加 -D stream.reduce.input.separator=,
