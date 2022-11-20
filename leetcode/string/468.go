package leetcode

import (
        "strconv"
)


// 468. Validate IP Address
//
// Given a string queryIP, return "IPv4" if IP is a valid IPv4 address, "IPv6" if IP is a valid IPv6 address 
// or "Neither" if IP is not a correct IP of any type.
// A valid IPv4 address is an IP in the form "x1.x2.x3.x4" where 0 <= xi <= 255 and xi cannot contain 
// leading zeros. For example, "192.168.1.1" and "192.168.1.0" are valid IPv4 addresses 
// while "192.168.01.1", "192.168.1.00", and "192.168@1.1" are invalid IPv4 addresses.
// A valid IPv6 address is an IP in the form "x1:x2:x3:x4:x5:x6:x7:x8" where:
// * 1 <= xi.length <= 4
// * xi is a hexadecimal string which may contain digits, lowercase English letter ('a' to 'f') 
//   and upper-case English letters ('A' to 'F').
// * Leading zeros are allowed in xi.
// * For example, "2001:0db8:85a3:0000:0000:8a2e:0370:7334" and "2001:db8:85a3:0:0:8A2E:0370:7334" 
// are valid IPv6 addresses, while "2001:0db8:85a3::8A2E:037j:7334" and "02001:0db8:85a3:0000:0000:8a2e:0370:7334" 
// are invalid IPv6 addresses.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Validate IP Address.
// Memory Usage: 2 MB, less than 83.33% of Go online submissions for Validate IP Address.
//
func validateIPAddress(queryIP string) string {
    if validateIPv4(queryIP) {
        return "IPv4"
    }
    if validateIPv6(queryIP) {
        return "IPv6"
    }
    return "Neither"
}

func validateIPv4(queryIP string) bool {
    if len(queryIP) < 7 || len(queryIP) > 15 {
        return false
    }
    parsed, low, high := 0, 0, 0
    for high <= len(queryIP) {
        if parsed == 4 {
            return false
        }
        if high == len(queryIP) || queryIP[high] == '.' {
            if high - low > 3 || high == low {
                return false
            }
            if high - low > 1 && queryIP[low] == '0' {
                return false
            }
            num, _ := strconv.Atoi(queryIP[low:high])
            if num > 255 {
                return false
            }
            parsed++
            low = high+1
        } else if queryIP[high] < '0' || queryIP[high] > '9' {
            return false
        }
        high++
    }
    if parsed != 4 {
        return false
    }
    return true
}

func validateIPv6(queryIP string) bool {
    if len(queryIP) < 15 || len(queryIP) > 39 {
        return false
    }
    parsed, low, high := 0, 0, 0
    for high <= len(queryIP) {
        if parsed == 8 {
            return false
        }
        if high == len(queryIP) || queryIP[high] == ':' {
            if high == low || high - low > 4 {
                return false
            }
            parsed++
            low = high+1
        } else if !(queryIP[high] >= 'a' && queryIP[high] <= 'f') && !(
                queryIP[high] >= 'A' && queryIP[high] <= 'F') && !(queryIP[high] >= '0' && queryIP[high] <= '9') {
            return false
        }
        high++
    }
    if parsed != 8 {
        return false
    }
    return true
}
