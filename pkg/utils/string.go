package utils

// Substr 截取字符串
// start: 正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
//        负数 - 在从字符串结尾的指定位置开始
// length: 正数 - 从 start 参数所在的位置返回
//         负数 - 从字符串末端返回
func Substr(str string, start, length int) string {
	if str == "" || length == 0 {
		return ""
	}

	runeStr := []rune(str)
	strLen := len(runeStr)

	if start < 0 {
		start = strLen + start
	}
	if start > strLen {
		start = strLen
	}
	end := start + length
	if end > strLen {
		end = strLen
	}
	if length < 0 {
		end = strLen + length
	}
	if start > end {
		start, end = end, start
	}

	return string(runeStr[start:end])
}
