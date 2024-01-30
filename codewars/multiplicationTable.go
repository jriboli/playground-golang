package main

// var _ = Describe("Sample Tests", func() {
// 	It("Sample Tests", func() {
// 		Expect(MultiplicationTable(1)).To(Equal([][]int{{1}}))
// 		Expect(MultiplicationTable(2)).To(Equal([][]int{{1, 2}, {2, 4}}))
// 		Expect(MultiplicationTable(3)).To(Equal([][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}))
// 	})
// })

func MultiplicationTable(size int) [][]int {
	// Implement me! :)
	var results [][]int
	
	for i:=1; i<=size; i++ {
		var singleArr []int

		for j:=1; j<=size; j++ {
			singleArr = append(singleArr, i*j)
		}

		results = append(results, singleArr)
	}

	
	return results
  }