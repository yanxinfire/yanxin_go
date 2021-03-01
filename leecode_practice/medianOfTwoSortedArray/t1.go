package medianOfTwoSortedArray

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	flag := 0
	if (m+n)%2 != 0 {
		flag = 1
	}
	mdmin := 0
	mdmax := 0
	if n == 0 && m != 0 {
		if m == 1 {
			mdmax = nums2[0]
		} else {
			mdmin = nums2[m/2-1]
			mdmax = nums2[m/2]
		}
	}
	if m == 0 && n != 0 {
		if n == 1 {
			mdmax = nums1[0]
		} else {
			mdmin = nums1[n/2-1]
			mdmax = nums1[n/2]
		}
	}
	for x, y := 0, 0; x < n; x++ {
		for y < m {
			if nums2[y] < nums1[x]{
				if x+y == (m+n)/2-1{
					mdmax=nums2[y+1]
					return res(mdmin,mdmax,flag)
				}
				if y==m-1{
					if x+y == (m+n)/2-1 {
						mdmax = nums2[y]
					}else {
						mdmin = nums1[(m+n)/2-y]
						mdmax = nums1[(m+n)/2-y+1]
					}
					fmt.Println("A")
					return res(mdmin,mdmax,flag)
				}
				mdmin = nums2[y]
				y++
				continue
			}
			if x == n-1 {
				if x+y == (m+n)/2-1 {
					mdmin = nums1[x]
					mdmax = nums2[y]
				} else {
					mdmin = nums2[(m+n)/2-x]
					mdmax = nums2[(m+n)/2-x+1]
				}
				return res(mdmin, mdmax, flag)
			}
			if x+y == (m+n)/2-1 {
				if nums2[y] > nums1[x] && nums2[y] < nums1[x+1] {
					mdmax = nums2[y]
				} else {
					mdmax = nums1[x+1]
				}
				return res(mdmin, mdmax, flag)
			}
			//fmt.Println(nums2[y],nums1[x],nums1[x+1])
			if nums2[y] > nums1[x] && nums2[y] < nums1[x+1] || nums2[y]<nums1[x]{
				mdmin = nums2[y]
				y++
			} else {
				if x==0{
					x=-1
				}
				break
			}
			if y == m-1 {
				if nums2[y] > nums1[x] && nums2[y] < nums1[x+1] {
					if x+y == (m+n)/2-1 {
						mdmin = nums2[y-1]
						mdmax = nums2[y]
					} else {
						mdmin = nums1[(m+n)/2-y]
						mdmax = nums1[(m+n)/2-y+1]
					}
					return res(mdmin, mdmax, flag)
				} else {
					if x+y == (m+n)/2-1 {
						mdmin = nums2[y-1]
						mdmax = nums1[x+1]
						return res(mdmin, mdmax, flag)
					}else {
						break
					}
				}
			}
		}
	}
	return res(mdmin, mdmax, flag)
}

func res(mdmin, mdmax, flag int) float64 {
	if flag == 0 {
		return float64(mdmin+mdmax) / 2
	}
	return float64(mdmax)
}
