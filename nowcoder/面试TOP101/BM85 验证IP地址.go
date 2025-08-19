package nowcoder

import (
	"strconv"
	"strings"
)

// https://www.nowcoder.com/practice/55fb3c68d08d46119f76ae2df7566880

func solveBM85(IP string) string {
	isIPV4 := func(ip string) bool {
		strs := strings.Split(ip, ".")
		if len(strs) != 4 {
			return false
		}

		for i := range strs {
			if len(strs[i]) == 0 {
				return false
			}

			if len(strs[i]) < 0 || len(strs[i]) > 3 || (strs[i][0] == '0' && len(strs[i]) != 1) {
				return false
			}

			for _, d := range strs[i] {
				if d < '0' || d > '9' {
					return false
				}
			}

			num, _ := strconv.Atoi(strs[i])
			if num < 0 || num > 255 {
				return false
			}
		}
		return true
	}

	isIPV6 := func(ip string) bool {
		strs := strings.Split(ip, ":")
		if len(strs) != 8 {
			return false
		}

		for i := range strs {

			if len(strs[i]) == 0 || len(strs[i]) > 4 {
				return false
			}

			for _, d := range strs[i] {
				if (d < '0' || d > '9') && (d < 'a' || d > 'f') && (d < 'A' || d > 'F') {
					return false
				}
			}
		}
		return true
	}

	if len(IP) == 0 {
		return "Neither"
	}

	if isIPV4(IP) {
		return "IPv4"
	}

	if isIPV6(IP) {
		return "IPv6"
	}

	return "Neither"
}
