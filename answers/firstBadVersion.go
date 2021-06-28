package answers

import "sort"

//https://leetcode-cn.com/problems/first-bad-version/
//278. 第一个错误的版本
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func FirstBadVersion(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
