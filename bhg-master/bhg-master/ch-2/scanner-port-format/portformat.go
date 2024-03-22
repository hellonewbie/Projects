package portformat

import (
	"errors"
	"strconv"
	"strings"
)

const (
	porterrmsg = "Invalid port specification"
)

func dashSplit(sp string, ports *[]int) error {
	//以-为分隔符进行分割，得到的结构以字符串的形式存入字符串切片中
	dp := strings.Split(sp, "-")
	if len(dp) != 2 {
		return errors.New(porterrmsg)
	}
	//将开始的端口号的形式从字符串转换为整型数
	start, err := strconv.Atoi(dp[0])
	if err != nil {
		return errors.New(porterrmsg)
	}
	end, err := strconv.Atoi(dp[1])
	if err != nil {
		return errors.New(porterrmsg)
	}
	//确保端口范围在有效范围里面
	if start > end || start < 1 || end > 65535 {
		return errors.New(porterrmsg)
	}
	for ; start <= end; start++ {
		*ports = append(*ports, start)
	}
	return nil
}

func convertAndAddPort(p string, ports *[]int) error {
	i, err := strconv.Atoi(p)
	if err != nil {
		return errors.New(porterrmsg)
	}
	if i < 1 || i > 65535 {
		return errors.New(porterrmsg)
	}
	*ports = append(*ports, i)
	return nil
}

// Parse turns a string of ports separated by '-' or ',' and returns a slice of Ints.
func Parse(s string) ([]int, error) {
	//创建一个空的整数切片，用于存储解析后的端口。
	ports := []int{}
	//如果字符串包含逗号和破折号，表示可能包含多个端口和端口范围。
	if strings.Contains(s, ",") && strings.Contains(s, "-") {
		sp := strings.Split(s, ",")
		for _, p := range sp {
			if strings.Contains(p, "-") {
				if err := dashSplit(p, &ports); err != nil {
					return ports, err
				}
			} else {
				if err := convertAndAddPort(p, &ports); err != nil {
					return ports, err
				}
			}
		}
	} else if strings.Contains(s, ",") {
		sp := strings.Split(s, ",")
		for _, p := range sp {
			convertAndAddPort(p, &ports)
		}
	} else if strings.Contains(s, "-") {
		if err := dashSplit(s, &ports); err != nil {
			return ports, err
		}
	} else {
		if err := convertAndAddPort(s, &ports); err != nil {
			return ports, err
		}
	}
	return ports, nil
}
