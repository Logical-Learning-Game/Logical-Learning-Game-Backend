package utility

import "github.com/lib/pq"

func PqInt32ArrayToIntSlice(pqInt32Array pq.Int32Array) []int {
	intSlice := make([]int, len(pqInt32Array))

	for i := range pqInt32Array {
		intSlice[i] = int(pqInt32Array[i])
	}

	return intSlice
}

func IntSliceToPqInt32Array(s []int) pq.Int32Array {
	pqInt32Array := make(pq.Int32Array, len(s))

	for i := range s {
		pqInt32Array[i] = int32(s[i])
	}

	return pqInt32Array
}
