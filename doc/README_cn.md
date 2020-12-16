# sacgo

一个go程序解析SAC(Seismic Analysis Code)文件

特性
========

- 跨平台
- 上手简易

例子
=======

- powershell

``` powershell
.\sacgo.exe -i test.SAC -h Delta        # 打印Delta值
.\sacgo.exe -i test.SAC -o test.ASC     # 生成文本文件
``` 

- batch

``` batch
sacgo.exe -i test.SAC -h Delta          # 打印Delta值
sacgo.exe -i test.SAC -o test.ASC       # 生成文本文件
``` 

- shell

``` shell
sacgo -i test.SAC -h Delta              // 打印Delta值
sacgo -i test.SAC -o test.ASC           // 生成文本文件
``` 