package zdpgo_clearcode

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// ClearCode 清除代码中的注释，空行
func ClearCode(filePath string) (string, error) {
	// 读取文件
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// 替换注释
	reg := regexp.MustCompile(`#.*`)
	result := reg.ReplaceAllString(string(fileContent), "")

	// 替换文档字符串
	reg = regexp.MustCompile(`'''[\s\S]*?'''`)
	result = reg.ReplaceAllString(result, "")

	reg = regexp.MustCompile(`"""[\s\S]*?"""`)
	result = reg.ReplaceAllString(result, "")

	// 替换//类型注释
	reg = regexp.MustCompile(`\s//.*`)
	result = reg.ReplaceAllString(result, "")

	// 替换代码行后面的 // 注释
	reg = regexp.MustCompile(`;\s*//.*`)
	result = reg.ReplaceAllString(result, ";")

	// 替换多行注释
	reg = regexp.MustCompile(`(\/\/.*$)|(\/\*(.|\s)*?\*\/)`)
	result = reg.ReplaceAllString(result, "")

	// 替换开头空行
	reg = regexp.MustCompile(`^\s*\n`)
	result = reg.ReplaceAllString(result, "")

	// 替换末尾空行
	reg = regexp.MustCompile(`\s*$`)
	result = reg.ReplaceAllString(result, "")

	// 替换中间空行
	reg = regexp.MustCompile(`\s*\n+`)
	result = reg.ReplaceAllString(result, "\n")
	strings.TrimSpace(result)

	// 返回
	return result, nil
}
