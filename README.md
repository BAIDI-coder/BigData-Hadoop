# BigData-Hadoop
## Streaming doc
### Purpose
- 计算城市薪资平均值的hadoop streaming script (by golang)

### Notice
- 架构：X86_64/amd64
- Hadoop 版本：3.2.1
- separator指定为西文逗号','

### Start
1. 获取
``` shell
	wget http://49.234.85.223/download/mapReducer.tar.gz
```
2. 解压
``` shell
	tar -zxvf mapReducer.tar.gz
```
3. 文件结构
``` shell
mapReducer
├── calAverage	/*reducer*/
├── cityMapper	/*mapper*/
├── debug.sh
├── input.txt	/*debug input*/
├── output
└── README.md

0 directories, 6 files
```
### Example
- 修改权限
``` shell
sudo chmod -R ugo+x ./mapReducer
```
- 调试
``` shell
cd ./mapReducer
./debug.sh
ls
cat output
```
- 使用样例(仅参考，实际文件路径依据情况而定)
1. 上传input.txt到HDFS
``` shell
hdfs dfs -put ./input.txt input
```
2. 执行如下streaming指令
``` shell
mapred streaming \
-D stream.map.output.field.separator=, \
-D stream.reduce.input.field.separator=, \
-files cityMapper,calAverage \
-input /user/hadoop/input/input.txt \
-output output \
-mapper cityMapper \
-reducer calAverage
```
3. 查看结果
``` shell
hdfs dfs -cat output/*
```

### Repository
- https://github.com/BAIDI-coder/BigData-Hadoop
