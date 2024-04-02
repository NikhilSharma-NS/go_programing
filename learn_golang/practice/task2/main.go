package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
A website domain "api.sailpoint.com" consists of various subdomains. At the top level, we have "com", at the next level, we have "sailpoint.com" and at the lowest level, "api.sailpoint.com". When we visit a
domain like "api.sailpoint.com", we will also visit the parent domains "sailpoint.com" and "com" implicitly.

A count-paired domain is a domain that has one of the two formats "rep d1.d2.d3" or "rep d1.d2" where rep is the number of visits to the domain and d1.d2.d3 is the domain itself.

For example, "25 api.sailpoint.com" is a count-paired domain that indicates that api.sailpoint.com was visited 25 times.
Given an array of count-paired domains cpdomains, return an array of the count-paired domains of each subdomain in the input. You may return the answer in any order.

Input: starts = ["25 api.sailpoint.com"]
Output: ["25 sailpoint.com","25 api.sailpoint.com","25 com"]

Input: ["25 api,sailpoint,com", "175 sailpoint.com","175 com","100 blog.api.sailpoint.com"]
Output: ["475 com", "300 sailpoint.com", "125 api.sailpoint.com","100 blog.api.sailpoint.com"]
*/
func main() {
	domains := []string{"25 api.sailpoint.com", "175 sailpoint.com", "175 com", "100 blog.api.sailpoint.com"}
	results := subdomainVisits(domains)
	for _, r := range results {
		fmt.Println(r)
	}
}

func subdomainVisits(ds []string) (out []string) {

	m := make(map[string]string)

	for _, v := range ds {
		g := strings.Split(v, " ")
		m[g[1]] = g[0]
		//fmt.Println(i)
	}

	for key, v1 := range m {
		v1_total, _ := strconv.ParseInt(v1, 10, 64)
		for key1, v := range m {
			if strings.Contains(key1, key) && !(key1 == key) {
				v_total, _ := strconv.ParseInt(v, 10, 64)
				v1_total = v1_total + v_total
				//fmt.Println(key, key1, v)
			}
		}
		out = append(out, strconv.FormatInt(v1_total, 10)+" "+key)

	}

	return out

}
