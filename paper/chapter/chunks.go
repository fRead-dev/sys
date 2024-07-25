package chapter

import "math"

//###########################################################//

const maxChunkSize = 63 * 1024 // 60 КБ в байтах

func ChunkSize(totalSize int) (int, int) {
	if totalSize <= 0 {
		return 0, 0
	}

	if totalSize < maxChunkSize {
		return totalSize, 1
	}

	numChunks := 0
	chunkSize := maxChunkSize
	for totalSize%chunkSize < chunkSize-100 {
		numChunks = int(math.Ceil(float64(totalSize) / float64(chunkSize)))

		chunkSize -= 10
	}

	return chunkSize, numChunks
}
