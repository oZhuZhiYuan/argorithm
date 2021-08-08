/* go test -v -run KthLargest
 */
package testcase

import (
	pro "argo/problems"
	util "argo/utils"
	"testing"
)

func TestKthLargest(t *testing.T) {
	kthLargest := pro.KthConstructor(3, []int{4, 5, 8, 2})
	testcase := []int{3, 5, 10, 9, 4}
	expect := []int{4, 5, 5, 8, 8}
	ret := []int{}
	for _, v := range testcase {
		ret = append(ret, kthLargest.Add(v))
	}
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test ReverseList fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test ReverseList success\nexpect: %d\nresult: %d", expect, ret)
}

func TestKthLargest1(t *testing.T) {
	kthLargest := pro.KthConstructor(2, []int{0})
	testcase := []int{-1, 1, -2, -4, 3}
	expect := []int{-1, 0, 0, 0, 1}
	ret := []int{}
	for _, v := range testcase {
		ret = append(ret, kthLargest.Add(v))
	}
	if !util.IntSliceCompare(expect, ret) {
		t.Fatalf("test ReverseList fail\nexpect: %d\nresult: %d", expect, ret)
	}

	t.Logf("test ReverseList success\nexpect: %d\nresult: %d", expect, ret)
}
