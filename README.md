# zdpgo_clearcode

清除代码中的注释，空行等内容，得到一个干净的代码

## 版本历史

- v0.1.0 新增：移除多余注释，空行
- v0.1.1 新增：切割代码字符串
- v0.1.2 新增：清除 Python 中的 main 代码
- v0.1.3 优化：切割代码字符串时清除空字符串
- v0.1.4 2022/08/08 优化：清除空格优化
- v0.1.6 2022/08/18 BUG修复：解决Python三引号字符串被置换为空的BUG
- v0.1.7 2022/08/19 新增：代码格式化功能

## 使用示例

### 移除各语言注释空行

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_clearcode"
)

func main() {
	filePathList := []string{
		"examples/test_data/level1_1.py",
		"examples/test_data/level1_1.java",
		"examples/test_data/level1_1.php",
		"examples/test_data/level1_1.c",
		"examples/test_data/level1_1.cpp",
		"examples/test_data/level1_2.py",
		"examples/test_data/level1_2.java",
		"examples/test_data/level1_2.php",
		"examples/test_data/level1_2.c",
		"examples/test_data/level1_2.cpp",
	}

	// 清除代码
	for _, filePath := range filePathList {
		fmt.Println(filePath)
		result, err := zdpgo_clearcode.ClearCode(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		fmt.Println("========================")
	}
}
```
