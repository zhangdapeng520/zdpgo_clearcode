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

	// 返回
	return result, nil
}

// SplitCode 按照指定的切割字符切割代码字符串，并移除包含在removeStrArr中的字符串
func SplitCode(codeStr string, splitStr string, removeStrArr []string) []string {
	// 切割字符串
	codeArr := strings.Split(codeStr, splitStr)

	// 移除字符串
	var result []string
	for _, code := range codeArr {
		// 移除空格，避免干扰判断
		code = strings.TrimSpace(code)

		// 判断是否需要移除
		isRemove := false
		for _, rm := range removeStrArr {
			if code == rm { // 相等移除
				isRemove = true
				break
			} else if strings.HasPrefix(rm, "^") && strings.HasPrefix(code, rm[1:]) { // 以指定前缀开头移除
				isRemove = true
				break
			}
		}

		// 如果不需要移除，则追加
		if !isRemove {
			result = append(result, code)
		}
	}

	// 返回
	return result
}