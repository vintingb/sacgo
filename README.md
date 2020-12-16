# sacgo

a go program read SAC(Seismic Analysis Code) flie

Features
========

- Cross-platform
- Simple

Document
========
[中文](doc/README_cn.md)

Examples
=======

- powershell

``` powershell
.\sacgo.exe -i test.SAC -h Delta        # Show the value of Delta
.\sacgo.exe -i test.SAC -o test.ASC     # Generate text file
``` 

- batch

``` batch
sacgo.exe -i test.SAC -h Delta          # Show the value of Delta
sacgo.exe -i test.SAC -o test.ASC       # Generate text file
``` 

- shell

``` shell
sacgo -i test.SAC -h Delta              // Show the value of Delta
sacgo -i test.SAC -o test.ASC           // Generate text file
``` 