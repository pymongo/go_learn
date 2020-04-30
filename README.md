# Go语言研究——HTTP客户端

## Go语言入门方法论：

1. 学习如何构建项目以及管理第三方库
2. 写一个http client/server处理json数据
3. 做一道leetcode题
4. 解决一个多线程问题(如哲学家进餐问题)

## Go测试与benchmark

实事求是的说，我更喜欢Rust给一个函数加上`#[test]`宏就能点绿色播放图标直接运行的简便。


## launchctl

查看当前所有资源的最大限制`launchctl limit`第一列是软件限制，第二列是硬件限制

```
localhost:~ w$ launchctl limit
	cpu         unlimited      unlimited      
	filesize    unlimited      unlimited      
	data        unlimited      unlimited      
	stack       8388608        67104768       
	core        0              unlimited      
	rss         unlimited      unlimited      
	memlock     unlimited      unlimited      
	maxproc     709            1064           
	maxfiles    1024           10240 
```
