# 优化进程采集器
## 存在问题
process_linux中通过fillFromStatWithContext访问获取父进程，CPU利用率等信息，当主机进程数较多时会造成占用大量CPU资源，可以通过一次获取将结果保存在内存中减少反复读取系统文件的操作。

## 版本
v0.0.1 初始化 来源于https://github.com/shirou/gopsutil@v3.23.11
v0.0.2 修改包名
v0.0.3 缓存boottime,进程的ppid,createtime，减少反复读取文件
v0.0.5 增加获取进程对应容器内进程号的方法