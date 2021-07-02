# progbar
progbar是一个基于Golang的进度条第三方包。 progbar is a package of license terms based on Golang.

操作简单，您只需要像包传两个参数，1.当前数值(int) 和 2.总数值(int)，progbar会根据您传入的两个参数进行计算，并将格式化后的进度条和百分比以string类型返回给调用程序。
Simple, you only need to pass two like packets, 1. the current value (int) and 2. the total parameter value (int, progbar will calculate your two parameters, the operation bar after the specific operation and the string type return to the call program.


```go
package main

import (
	"fmt"
	"progbar"
)

func main() {
	val, err := progbar.Display(83, 128)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(val)
}
```
输出结果/output：
```
[################################                        ] 64 %    (83/128)
```

## 模拟程序
```go
func main() {
	fileMax := 300

	//模拟大数据清洗过程,变量i递增代表已清洗的进度
	for i := 0; i <= fileMax; i++ {

		//为不影响程序性能,采用一下方式调用进度条
		if i%3 == 0 {
			vle, err := progbar.Display(i, fileMax)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Print(vle)
		}

		time.Sleep(30 * time.Millisecond)
	}
}
```


<font color=red>**注：**本软件软件包仅仅是作者学习golang时编写的玩具</font>
