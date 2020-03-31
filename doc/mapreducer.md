## Purpose
- 计算城市薪资平均值的hadoop streaming script (by golang)

## Notice
- 架构：X86_64/amd64
- Hadoop 版本：3.2.1
- separator指定为西文逗号','

## Start
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
mapReducer/
├── calAverage	//reducer
├── cityMapper	//mapper
├── debug.sh
└── README.md

0 directories, 4 files
```
## Example
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
``` shell
hadoop jar /usr/local/hadoop/share/hadoop/tools/lib/hadoop-streaming-3.2.1.jar \
  -D stream.map.output.field.separator=, \
  -input ~/Desktop/input.txt \
  -output ~/Desktop/output \
  -mapper ~/cityMapper -file ~/cityMapper \
  -reducer ~/calAverage -file ~/calAverage\
```
## Repository
- https://github.com/BAIDI-coder/BigData-Hadoop
