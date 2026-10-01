package main

import "fmt"

/*
Top K Frequent Elements

Pattern: Frequency + Top K → Hash Map + Bucket Sort.

Approach:
The dumb way is counting the frequencies and then sorting everything
by frequency. That works, but sorting makes it O(n log n).

Instead, first use a hash map to count how many times each number appears.

Then use buckets where the index is the frequency:
bucket[1] → numbers that appear once
bucket[2] → numbers that appear twice
bucket[3] → numbers that appear three times

Since no number can appear more than n times, we only need n + 1 buckets.

Then walk the buckets backwards, starting from the highest frequency,
and keep adding numbers until we have k of them.

The main thing to remember: when we need the Top K based on frequency,
think HASH MAP + BUCKET SORT.

Time: O(n)
Space: O(n)
*/

func topKFrequent(nums []int, k int) []int {
	// Okay, so let me first understand what I'm actually looking for.
	// I need the k numbers that show up the most, not the k biggest numbers.
	// Like [1,1,1,2,2,3], k=2 -> 1 appears 3 times, 2 appears 2 times,
	// so [1,2] is what I want. Basically, I'm gonna count every number
	// first, then I need some way to grab the k biggest counts.

	// So first I'm gonna maintain a map where number -> frequency,
	// and I'll loop over the whole nums slice, just increasing the count
	// every time I see the same number.
	frequency := make(map[int]int)

	for _, num := range nums {
		frequency[num]++
	}

	// Alright, now I have the frequency of everything. I could sort these
	// by frequency and take the first k, but that's O(n log n), and we don't
	// really need to sort. A number can appear at most len(nums) times,
	// so I'm gonna make buckets where the index itself means the frequency.
	//
	// bucket[1] means "appeared once", bucket[2] means "appeared twice", etc.
	// Since the frequency can go all the way from 1 to len(nums), I need
	// indexes 1 through len(nums). Index 0 isn't used, so I need one extra
	// slot. That's why it's len(nums)+1.
	//
	// Like if n = 4, possible frequencies are 1, 2, 3, 4,
	// so I need 5 slots: index 0, 1, 2, 3, 4.
	buckets := make([][]int, len(nums)+1)

	// Now I'm gonna loop over the frequency map and put each number
	// into the bucket matching its count. So if 1 appeared 3 times,
	// it goes into bucket[3]. This way the bucket position already
	// tells me how frequent the number is, no sorting needed.
	for num, count := range frequency {
		buckets[count] = append(buckets[count], num)
	}

	// Okay, now I just need k numbers. I'm gonna keep the result here,
	// then start from the biggest bucket and walk backwards because
	// obviously I want the highest frequencies first.
	result := make([]int, 0, k)

	for count := len(buckets) - 1; count >= 0 && len(result) < k; count-- {
		// If nobody appeared this many times, there's nothing here,
		// so uhh, just skip this bucket and keep going down.
		if len(buckets[count]) == 0 {
			continue
		}

		// Okay, this bucket actually has numbers. They all appeared
		// 'count' times, so I'm gonna add them to the answer.
		// If I reach k in the middle of this bucket, I'm done anyway.
		for _, num := range buckets[count] {
			result = append(result, num)

			if len(result) == k {
				break
			}
		}
	}

	// Alright, we got the k most frequent numbers.
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	result := topKFrequent(nums, k)

	fmt.Println(result)
}
