package leetcode

import "strings"

// 93. Restore IP Addresses
//
// A valid IP address consists of exactly four integers separated by single dots. 
// Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.
// For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" 
// and "192.168@1.1" are invalid IP addresses.
// Given a string s containing only digits, return all possible valid IP addresses 
// that can be formed by inserting dots into s. 
// You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.
//
// Runtime: 2 ms, faster than 42.57% of Go online submissions for Restore IP Addresses.
// Memory Usage: 2 MB, less than 86.14% of Go online submissions for Restore IP Addresses.
//
func restoreIpAddresses(s string) []string {
    if len(s) < 4 || len(s) > 12 {
        return []string{}
    }

    result := []string{}
    ip := []string{}

    solveIpDFS(&s, &ip, 0, &result)
    return result
}

func solveIpDFS(s *string, ip *[]string, offset int, ret *[]string){
    if len((*ip)) == 4 && offset == len(*s) {
        ipLen := 3 + len((*ip)[0]) + len((*ip)[1]) + len((*ip)[2]) + len((*ip)[3])
        sb := new(strings.Builder)
        sb.Grow(ipLen)
        sb.WriteString((*ip)[0])
        sb.WriteByte('.')
        sb.WriteString((*ip)[1])
        sb.WriteByte('.')
        sb.WriteString((*ip)[2])
        sb.WriteByte('.')
        sb.WriteString((*ip)[3])
        *ret = append(*ret, sb.String())
        return
    }

    if offset > len(*s) || len(*ip) > 4 {
        return
    }

    for i := 0; i < 3; i++ {
        if offset + i >= len(*s) {
            break
        }
        seg := (*s)[offset:offset+i+1]
        if (i == 0 && seg >= "0" && seg <= "9") || (i == 1 && seg >= "10" && seg <= "99") || (i == 2 && seg >= "100" && seg <= "255") {
            *ip = append(*ip, seg)
            solveIpDFS(s, ip, offset+i+1, ret)
            *ip = (*ip)[:len(*ip)-1]
        }
    }
}
