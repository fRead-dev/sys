package sys

import (
	"bytes"
	"testing"
)

func TestCompress(t *testing.T) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data"),
	}

	compressedData, err := data.Compress()
	if err != nil {
		t.Fatalf("Failed to compress data: %v", err)
	}

	if !compressedData.IsCompress {
		t.Errorf("Expected IsCompress to be true, got false")
	}

	if bytes.Equal(data.Data, compressedData.Data) {
		t.Errorf("Expected compressed data to be different from original data")
	}
}

func TestDecompress(t *testing.T) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data"),
	}
	compressedData, err := data.Compress()
	if err != nil {
		t.Fatalf("Failed to compress data: %v", err)
	}

	decompressedData, err := compressedData.Decompress()
	if err != nil {
		t.Fatalf("Failed to decompress data: %v", err)
	}

	if decompressedData.IsCompress {
		t.Errorf("Expected IsCompress to be false, got true")
	}

	if !bytes.Equal(decompressedData.Data, data.Data) {
		t.Errorf("Decompressed data does not match the expected value")
	}
}

func TestEncrypt(t *testing.T) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data"),
	}
	key := "test key"

	cData, err := data.Encrypt(key)
	if err != nil {
		t.Fatalf("Failed to Encrypt data: %v", err)
	}

	if !cData.IsCrypt {
		t.Errorf("Expected IsCrypt to be true, got false")
	}

	if bytes.Equal(data.Data, cData.Data) {
		t.Errorf("Expected Encrypt data to be different from original data")
	}
}

func TestDecrypt(t *testing.T) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data"),
	}
	key := "test key"

	cData, err := data.Encrypt(key)
	if err != nil {
		t.Fatalf("Failed to Encrypt data: %v", err)
	}

	deData, err := cData.Decrypt(key)
	if err != nil {
		t.Fatalf("Failed to Decrypt data: %v", err)
	}

	if deData.IsCompress {
		t.Errorf("Expected IsCrypt to be false, got true")
	}

	if !bytes.Equal(deData.Data, data.Data) {
		t.Errorf("Decrypt data does not match the expected value")
	}
}

func TestChunks(t *testing.T) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("this is a test of chunks"),
	}

	chunks, err := data.Chunks()
	if err != nil {
		t.Fatalf("Failed to get chunks: %v", err)
	}

	if len(chunks) == 0 {
		t.Errorf("Expected non-zero number of chunks")
	}

	// Validate each chunk
	for i, chunk := range chunks {
		if chunk.IsCompress != data.IsCompress {
			t.Errorf("Chunk %d has incorrect IsCompress value", i)
		}
		if len(chunk.Data) == 0 {
			t.Errorf("Chunk %d has empty data", i)
		}
	}
}

func TestChunksToData(t *testing.T) {
	// Подготовка данных для теста
	data1 := &DataObj{
		IsCompress: false,
		Data:       []byte("chunk1 "),
	}
	data2 := &DataObj{
		IsCompress: false,
		Data:       []byte("chunk2 "),
	}
	data3 := &DataObj{
		IsCompress: false,
		Data:       []byte("compressed chunk"), // Используйте корректные сжатые данные для вашего случая
	}

	// Создание объектов с данными
	chunks := []DataObj{*data1, *data2, *data3}

	// Запуск функции
	combinedData, err := ChunksToData(chunks)
	if err != nil {
		t.Fatalf("Failed to combine chunks: %v", err)
	}

	// Ожидаемые данные
	expectedData := []byte("chunk1 chunk2 compressed chunk") // Обновите в зависимости от того, что возвращает Decompress

	// Проверка результатов
	if !bytes.Equal(combinedData.Data, expectedData) {
		t.Errorf("Combined data does not match expected data. Got %s, want %s", combinedData.Data, expectedData)
	}
}

///////////////////////////////////////////////

func BenchmarkCompress(b *testing.B) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data for compression"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := data.Compress()
		if err != nil {
			b.Fatalf("Failed to compress data: %v", err)
		}
	}
}

func BenchmarkDecompress(b *testing.B) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data"),
	}
	compressedData, err := data.Compress()
	if err != nil {
		b.Fatalf("Failed to compress data: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := compressedData.Decompress()
		if err != nil {
			b.Fatalf("Failed to decompress data: %v", err)
		}
	}
}

func BenchmarkChunks(b *testing.B) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data to split into chunks"),
	}

	for i := 0; i < b.N; i++ {
		_, err := data.Chunks()
		if err != nil {
			b.Fatalf("Failed to get chunks: %v", err)
		}
	}
}

func BenchmarkChunksToData(b *testing.B) {
	var chunks []DataObj
	for i := 0; i < 100; i++ {
		data := &DataObj{
			IsCompress: false,
			Data:       []byte("test chunk data "),
		}
		chunks = append(chunks, *data)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ChunksToData(chunks)
		if err != nil {
			b.Fatalf("Failed to combine chunks: %v", err)
		}
	}
}
