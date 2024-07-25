package chapter

import "testing"

func TestChunkSize(t *testing.T) {
	tests := []struct {
		totalSize      int
		expectedSize   int
		expectedChunks int
	}{
		{0, 0, 0},
		{1, 1, 1},
		{60 * 1024, 60 * 1024, 1},
		{60*1024 + 1, 60*1024 + 1, 1},
		{120 * 1024, 60 * 1024, 2},
		{121 * 1024, 61 * 1024, 2},
		{122 * 1024, 61 * 1024, 2},
		{123 * 1024, 62 * 1024, 2},
		{2 * 1024 * 1024, 60 * 1024, 33},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			gotSize, gotChunks := ChunkSize(tt.totalSize)
			if gotChunks != tt.expectedChunks {
				t.Errorf("ChunkSize(%d) = (%d, %d :%d); want (%d, %d :%d)",
					tt.totalSize,
					gotSize, gotChunks, gotSize*gotChunks-tt.totalSize,
					tt.expectedSize, tt.expectedChunks, tt.expectedSize*tt.expectedChunks-tt.totalSize,
				)
			}
		})
	}
}

func BenchmarkChunkSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ChunkSize(1_000_000_000)
	}
}
