package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slices are dynamic arrays

	// 1. Creating a slice using make
	// Syntax: make([]type, len, cap)
	// cap is optional but good for performance if size is known
	nums := make([]int, 0, 10)
	// Alternative: nums := []int{}

	fmt.Printf("Initial capacity: %d\n", cap(nums))

	// 2. Adding elements to slice
	nums = append(nums, 1)
	nums = append(nums, 45)

	fmt.Println("Current slice:", nums)

	// 3. Copying a slice
	// Destination slice must calculate length based on source
	nums2 := make([]int, len(nums))
	copy(nums2, nums)

	fmt.Println("Copied slice (nums2):", nums2)

	// 4. Slicing a slice
	// Syntax: slice[start:end]
	// end is exclusive
	nums3 := nums[0:2]  // slice from index 1 to 2 (exclusive) 
	// we can also write it as 
	// nums3 := nums[:2]
	// nums3 := nums[0:] // slice from index 0 to end
	fmt.Println("Sliced slice (nums3):", nums3)
 
	// 5. Deleting from a slice
	// Syntax: append(slice[:index], slice[index+1:]...)
	nums = append(nums[:1], nums[2:]...)
	fmt.Println("Deleted slice (nums):", nums)

	// 6. Checking if a slice is empty
	if len(nums) == 0 {
		fmt.Println("Slice is empty")
	}

	// 7. checking arrays are equal using slices.Equal() , slices is a package in go
	fmt.Println(slices.Equal(nums2 , nums3))

	// 8. 2D slices
	 var slice_2d = [][]int{} // empty 2D slice
	 slice_2d = append(slice_2d, []int{1,2,3})
	 slice_2d = append(slice_2d, []int{4,5,6})
	 fmt.Println("2D slice:", slice_2d)

	 // 9. 3D slices
	 var slice_3d = [][][]int{} // empty 3D slice
	 slice_3d = append(slice_3d, [][]int{})
	 
}