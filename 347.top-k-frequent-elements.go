/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (64.88%)
 * Likes:    10898
 * Dislikes: 412
 * Total Accepted:    1.1M
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given an integer array nums and an integer k, return the k most frequent
 * elements. You may return the answer in any order.
 * 
 * 
 * Example 1:
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 * Example 2:
 * Input: nums = [1], k = 1
 * Output: [1]
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * k is in the range [1, the number of unique elements in the array].
 * It is guaranteed that the answer is unique.
 * 
 * 
 * 
 * Follow up: Your algorithm's time complexity must be better than O(n log n),
 * where n is the array's size.
 * 
 */
package leetcode

import "container/heap"
// @lc code=start
func topKFrequent(nums []int, k int) []int {
	// create a empty hashmap
	m := map[int]int{}
	// calculate the frequent counts of all the numers in nums
	for _, n := range nums {
		m[n]++
	} 
	//Use priorityQueue struct (heap)
	//Use a max Heap, push all Item to the max Heap
	q := PriorityQueue{}
	for key, count := range m {
		// add Item to heap 
		heap.Push(&q, &Item{key: key, count: count})
	}
	result := []int{}
	// use heap Pop k time root node, add the Item.key to a slice to present Top K Frequent Elements.
	for len(result) < k {
		item := heap.Pop(&q).(*Item)
		result = append(result, item.key)
	}
	return result
}
// https://www.sobyte.net/post/2022-07/go-heap/
// create a struct named Item
type Item struct{
	key int
	count int 
}
// create a Priority Queue impelment with heap interface and holds items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int  {
	return len(pq)
}

func (pq PriorityQueue) Less(i,j int) bool {
	// for default heap in golang is minHeap, so > make this a maxHeap
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i,j int){
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
	
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

// @lc code=end

