# 优化进程采集器
## 存在问题
process_linux中通过fillFromStatWithContext访问获取父进程，CPU利用率等信息，当主机进程数较多时会造成占用大量CPU资源，可以通过一次获取将结果保存在内存中减少反复读取系统文件的操作。

## 版本
0.0.1 初始化 来源于https://github.com/shirou/gopsutil@v3.23.11