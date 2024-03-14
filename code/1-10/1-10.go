package __10

//		执行耗时:241 ms,击败了5.02% 的Go用户
//		内存消耗:6.7 MB,击败了6.39% 的Go用户
//	 暴力循环 + hash
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	maxLen := 1
	for i := 0; i < len(s); i++ {
		// 生成一个hash表
		m := map[byte]bool{}

		curLen := 1
		m[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if m[s[j]] {
				curLen = j - i
				break
			}
			m[s[j]] = true
			curLen = j - i + 1
		}
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

//		执行耗时:37 ms,击败了9.67% 的Go用户
//		内存消耗:2.8 MB,击败了55.38% 的Go用户
//		滑动窗口 + hash
//	 难点在于如何删除已经出现过重复的字符
func LengthOfLongestSubstring1(s string) int {
	if s == "" {
		return 0
	}
	var flag1, flag2 int
	maxLen := 0
	// 生成一个hash表
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok {
			curLen := flag2 - flag1 + 1
			if curLen > maxLen {
				maxLen = curLen
			}
			flag1 = index + 1
			// 删除之前的
			for k, v := range m {
				if v < flag1 {
					delete(m, k)
				}
			}
		}
		m[s[i]] = i
		flag2 = i
	}
	curLen := flag2 - flag1 + 1
	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}

//		执行耗时:8 ms,击败了62.43% 的Go用户
//		内存消耗:2.5 MB,击败了77.43% 的Go用户
//	 不需要便利循环hash删除
func LengthOfLongestSubstring2(s string) int {
	if s == "" {
		return 0
	}

	m := make(map[byte]bool)
	var left, maxLen, curLen int
	for i := 0; i < len(s); i++ {
		curLen += 1
		for m[s[i]] {
			delete(m, s[left])
			left += 1
			curLen -= 1
		}
		if curLen > maxLen {
			maxLen = curLen
		}
		m[s[i]] = true
	}
	return maxLen
}

// LongestPalindrome 最长的回文子串
// eg: babad
// 动态规划
//
//	执行耗时:78 ms,击败了21.68% 的Go用户
//	内存消耗:6.8 MB,击败了20.30% 的Go用户
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	strLen := len(s)
	m := make([][]bool, strLen)
	for i := 0; i < strLen; i++ {
		m[i] = make([]bool, strLen)
	}
	maxLen := 1
	si := 0
	for curLen := 2; curLen <= strLen; curLen++ {
		for i := 0; i < strLen; i++ {
			// j - i + 1 = curLen
			j := curLen + i - 1
			// 超出边界，退出
			if j >= strLen {
				break
			}
			if s[i] != s[j] {
				m[i][j] = false
			} else {
				// 长度小于等于3的，都可以直接判定true
				if curLen <= 3 {
					m[i][j] = true
				} else {
					m[i][j] = m[i+1][j-1]
				}
			}
			if m[i][j] && curLen > maxLen {
				si = i
				maxLen = curLen
			}
		}
	}
	return s[si : si+maxLen]
}

// LongestPalindrome1 最长的回文子串
// eg: babad cbbd
// 中心扩散法
// 需要注意点
//  1. 当最长的为1的时候 比如abcd
//     执行耗时:5 ms,击败了60.28% 的Go用户
//     内存消耗:2.3 MB,击败了60.01% 的Go用户
func LongestPalindrome1(s string) string {
	length := len(s)
	if length == 0 {
		return s
	}
	var max int = 0
	var str string
	for cur := 0; cur < length; cur++ {
		left, right, llen, rlen := cur-1, cur+1, 0, 0
		// 1、先向左边走，当左边位置等于当前位置，继续向左边走，直到不相等（left=curr）
		for left > 0 && s[left] == s[cur] {
			left--
			//左移数+1
			llen++
		}
		// 2、然后向右边走，当右边位置等于当前位置，继续向右边走，直到不相等（right=curr）
		for right < length && s[right] == s[cur] {
			right++
			//右移数+1
			rlen++
		}
		// 3、判断左边位置和右边位置是否相等，相等则继续向两边走，直到不相等（left=right）
		for left >= 0 && right < length {
			if s[left] != s[right] {
				break
			}
			left--
			right++
			llen++
			rlen++
		}
		//长度大于最大的重新确定最长回文字符串
		if llen+rlen > max {
			max = llen + rlen
			str = s[cur-llen : cur+rlen+1]
		}
	}
	//默认第一个字符
	if max == 0 {
		str = string([]byte(s)[0])
	}
	return str
}

// Convert Z 字形排列
// eg: 输入：s = "PAYPALISHIRING", numRows = 3 输出："PAHNAPLSIIGYIR"
//
//	执行耗时:11 ms,击败了40.76% 的Go用户
//	内存消耗:7.2 MB,击败了26.47% 的Go用户
func Convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	d := make([]string, numRows)
	num := 1
	direction := -1
	for i := 0; i < len(s); i++ {
		num = num + direction
		d[num] += string(s[i])
		if num == 0 || num == numRows-1 {
			direction *= -1
		}
	}
	r := ""
	for _, i := range d {
		r += i
	}
	return r
}

// Reverse 整数反转
// eg: 123 321
// eg: 120 21
// eg: -12 -21
// eg: 0 0
//
//	执行耗时:4 ms,击败了23.79% 的Go用户
//	内存消耗:2 MB,击败了71.24% 的Go用户
func Reverse(x int) int {
	num := 0
	negative := x < 0
	if x < 0 {
		x *= -1
	}
	for x != 0 {
		e := x % 10
		x /= 10
		num = num*10 + e
		if num > 1<<31-1 {
			num = 0
		}
	}
	if negative {
		num = num * -1
	}
	return num
}

// MyAtoi 8. 字符串转换整数
//
//	--       空格       +/-     数字      其他
//
// start     start    sign    inNumber   end
// sign      end      end 	  inNumber   end
// inNumber  end      end 	  inNumber   end
// end  	 end      end     end        end
// 状态机
//
//	执行耗时:4 ms,击败了28.53% 的Go用户
//	内存消耗:2.1 MB,击败了51.30% 的Go用户
var (
	start    = 0
	sign     = 1
	inNumber = 2
	end      = 3

	space  = 0
	other  = 3
	status = [][]int{
		{start, sign, inNumber, end},
		{sign, end, inNumber, end},
		{end, end, inNumber, end},
		{end, end, end, end},
	}
)

func MyAtoi(s string) int {
	curStatus := start
	negative := false
	num := 0
	for i := 0; i < len(s); i++ {
		curStatus = getMyAtoiStatus(curStatus, s[i])
		if curStatus == start {
			continue
		} else if curStatus == sign {
			if s[i] == '-' {
				negative = true
			}
		} else if curStatus == inNumber {
			num = num*10 + int(s[i]-'0')
			if negative {
				if num > 1<<31 {
					return -1 << 31
				}
			} else {
				if num > 1<<31-1 {
					return 1<<31 - 1
				}
			}
		} else {
			break
		}
	}
	if negative {
		num *= -1
	}
	return num
}

func getMyAtoiStatus(lastStatus int, d byte) int {
	currentStatus := end
	if d == ' ' {
		currentStatus = space
	} else if d == '-' || d == '+' {
		currentStatus = sign
	} else if d >= '0' && d <= '9' {
		currentStatus = inNumber
	} else {
		currentStatus = other
	}
	return status[lastStatus][currentStatus]
}

// IsPalindrome 9. 回文数
//
//		 反转数字
//	  重点在于不能出现后面为0的数
//			执行耗时:16 ms,击败了42.68% 的Go用户
//			内存消耗:4.2 MB,击败了57.56% 的Go用户
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	num := 0
	for num < x {
		num = num*10 + x%10
		x /= 10
	}
	return x == num || x == num/10
}
