package day05

import (
	"fmt"
	"slices"
	"testing"
)

func TestSolution1(t *testing.T) {
	cases := map[string]int{
		"test_input11": 3,
	}

	for inputFile, expectedResult := range cases {
		result := Solution1(inputFile)
		if result != expectedResult {
			t.Fatalf("Solution1() = %d, expecting %d\n", result, expectedResult)
		} else {
			fmt.Printf("Solution1() = %d, OK\n", result)
		}
	}
}

// func TestSolution2(t *testing.T) {
// 	cases := map[string]int{
// 		"test_input21": 46,
// 	}
//
// 	for inputFile, expectedResult := range cases {
// 		result := Solution2(inputFile)
// 		if result != expectedResult {
// 			t.Fatalf("Solution2() = %d, expecting %d\n", result, expectedResult)
// 		} else {
// 			fmt.Printf("Solution2() = %d, OK\n", result)
// 		}
// 	}
// }

func TestInputRangeApplyMapping(t *testing.T) {
	inputRange := InputRange{start: 10, length: 11} // inclusive 10..20
	overlapCases := map[string][][][]int{
		"left outer unmapped": [][][]int{
			[][]int{{15, 100, 11}}, // MappingRange srcStart, destStart, length
			[][]int{{100, 6}},      // mapped InputRanges start, length
			[][]int{{10, 5}},       // unmapped InputRanges start, length
		},
		"right outer unmapped": [][][]int{
			[][]int{{5, 100, 11}},
			[][]int{{105, 6}},
			[][]int{{16, 5}},
		},
		"right outer unmapped, 1 char overlap": [][][]int{
			[][]int{{5, 100, 6}},
			[][]int{{105, 1}},
			[][]int{{11, 10}},
		},
		"only middle section mapped": [][][]int{
			[][]int{{14, 100, 6}},
			[][]int{{100, 6}},
			[][]int{{10, 4}, {20, 1}},
		},
		"only left outder unmapped, 1 char overlap": [][][]int{
			[][]int{{20, 100, 10}},
			[][]int{{100, 1}},
			[][]int{{10, 10}},
		},
		"whole range mapped": [][][]int{
			[][]int{{10, 100, 11}},
			[][]int{{100, 11}},
			[][]int{},
		},
	}

	for caseName, params := range overlapCases {
		mappingRangeParams := params[0][0]
		mappingRange := MappingRange{srcStart: mappingRangeParams[0], destStart: mappingRangeParams[1], length: mappingRangeParams[2]}

		var expectedMappedRanges []InputRange
		var expectedUnmappedRanges []InputRange

		for _, expectedMappedParams := range params[1] {
			expectedMappedRanges = append(expectedMappedRanges, InputRange{start: expectedMappedParams[0], length: expectedMappedParams[1]})
		}

		for _, expectedUnmappedParams := range params[2] {
			expectedUnmappedRanges = append(expectedUnmappedRanges, InputRange{start: expectedUnmappedParams[0], length: expectedUnmappedParams[1]})
		}

		mapped, unmapped, ok := inputRange.applyMapping(mappingRange)

		if !ok || len(mapped) != len(expectedMappedRanges) || len(unmapped) != len(expectedUnmappedRanges) {
			t.Fatalf("InputRange.applyMapping() %v, ok = %t, mapped= %v, unmapped= %v; FAIL\n", caseName, ok, mapped, unmapped)
		}

		for _, mappedRange := range mapped {
			if !slices.Contains(expectedMappedRanges, mappedRange) {
				t.Fatalf("InputRange.applyMapping() %v, unexpected mapped range %v; FAIL\n", caseName, mappedRange.describe())
			}
		}

		for _, unmappedRange := range unmapped {
			if !slices.Contains(expectedUnmappedRanges, unmappedRange) {
				t.Fatalf("InputRange.applyMapping() %v, unexpected unmapped range %v; FAIL\n", caseName, unmappedRange.describe())
			}
		}

		fmt.Printf("InputRange.applyMapping() %v - OK\n", caseName)
	}
}

func TestMappingRangeFindDest(t *testing.T) {
	rangeStr := MappingRange{srcStart: 10, destStart: 20, length: 10}
	cases := map[int]int{
		10: 20,
		20: 30,
		15: 25,
		9:  -1,
		31: -1,
	}

	for srcStart, expectedDest := range cases {
		result := rangeStr.findDest(srcStart)
		if result != expectedDest {
			t.Fatalf("Range.findDest() = %d, expecting %d\n", result, expectedDest)
		} else {
			fmt.Printf("Range.findDest() = %d, OK\n", result)
		}
	}
}
