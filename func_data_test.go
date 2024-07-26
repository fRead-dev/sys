package sys

import (
	"bytes"
	"strings"
	"testing"
)

var dataTest = DataObj{Data: []byte("test data")}
var keyTest = "test key"

func testCompress(t *testing.T, data *DataObj) *DataObj {
	compressedData, err := data.Compress()
	if err != nil {
		t.Fatalf("Failed to compress data: %v", err)
	}

	if !compressedData.IsCompress {
		t.Fatalf("Expected IsCompress to be true, got false")
	}

	if bytes.Equal(data.Data, compressedData.Data) {
		t.Fatalf("Expected compressed data to be different from original data")
	}

	return compressedData
}

func testDecompress(t *testing.T, data *DataObj) *DataObj {
	decompressedData, err := data.Decompress()
	if err != nil {
		t.Fatalf("Failed to decompress data: %v", err)
	}

	if decompressedData.IsCompress {
		t.Fatalf("Expected IsCompress to be false, got true")
	}

	if bytes.Equal(decompressedData.Data, data.Data) {
		t.Fatalf("Decompressed data does not match the expected value")
	}

	return decompressedData
}

func testEncrypt(t *testing.T, data *DataObj, key string) *DataObj {
	cData, err := data.Encrypt(key)
	if err != nil {
		t.Fatalf("Failed to Encrypt data: %v", err)
	}

	if !cData.IsCrypt {
		t.Errorf("Expected IsCrypt to be true, got false")
	}

	if bytes.Equal(cData.Data, data.Data) {
		t.Fatalf("Encrypt data does not match the expected value")
	}

	return cData
}

func testDecrypt(t *testing.T, data *DataObj, key string) *DataObj {
	deData, err := data.Decrypt(key)
	if err != nil {
		t.Fatalf("Failed to Decrypt data: %v", err)
	}

	if deData.IsCrypt {
		t.Errorf("Expected IsCrypt to be false, got true")
	}

	if bytes.Equal(deData.Data, data.Data) {
		t.Fatalf("Decrypt data does not match the expected value")
	}

	return deData
}

//###########################################################//

func TestCompress(t *testing.T) {
	testCompress(t, &dataTest)
}

func TestDecompress(t *testing.T) {
	compressedData := testCompress(t, &dataTest)
	decompressedData := testDecompress(t, compressedData)

	if !bytes.Equal(dataTest.Data, decompressedData.Data) {
		t.Errorf("Expected Compress data to be different from original data")
	}
}

func TestEncrypt(t *testing.T) {
	testEncrypt(t, &dataTest, keyTest)
}

func TestDecrypt(t *testing.T) {
	encryptData := testEncrypt(t, &dataTest, keyTest)
	decryptData := testDecrypt(t, encryptData, keyTest)

	if !bytes.Equal(dataTest.Data, decryptData.Data) {
		t.Errorf("Expected Crypt data to be different from original data")
	}
}

////

func TestChunks(t *testing.T) {
	data := &DataObj{
		Data: []byte("this is a test of chunks"),
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
		Data: []byte("chunk1 "),
	}
	data2 := &DataObj{
		Data: []byte("chunk2 "),
	}
	data3 := &DataObj{
		Data: []byte("chunk3"),
	}

	// Создание объектов с данными
	chunks := []DataObj{*data1, *data2, *data3}

	// Запуск функции
	combinedData, err := ChunksToData(chunks)
	if err != nil {
		t.Fatalf("Failed to combine chunks: %v", err)
	}

	// Ожидаемые данные
	expectedData := []byte("chunk1 chunk2 chunk3") // Обновите в зависимости от того, что возвращает Decompress

	// Проверка результатов
	if !bytes.Equal(combinedData.Data, expectedData) {
		t.Errorf("Combined data does not match expected data. Got %s, want %s", combinedData.Data, expectedData)
	}
}

///////////////////////////////////////////////

func BenchmarkCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := dataTest.Compress()
		if err != nil {
			b.Fatalf("Failed to compress data: %v", err)
		}
	}
}

func BenchmarkDecompress(b *testing.B) {
	compressedData, err := dataTest.Compress()
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
		Data:       []byte(strings.Repeat(string(dataTest.Data), 1_000_000)),
	}

	b.ResetTimer()
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
		data := &dataTest
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
