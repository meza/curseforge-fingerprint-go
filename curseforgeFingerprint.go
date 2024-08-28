package curseforgeFingerprint

import (
	"bufio"
	"github.com/aviddiviner/go-murmur"
	"os"
	"sync"
)

const chunkSize = 1024

func GetFingerprintFor(filePath string) uint32 {
	file, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	result := processFileInChunks(file)
	return getByteArrayHash(result)
}

func processFileInChunks(file *os.File) []byte {
	var wg sync.WaitGroup
	reader := bufio.NewReader(file)

	chunkOrder := 0
	chunkMap := make(map[int][]byte)
	var mu sync.Mutex

	for {
		chunk := make([]byte, chunkSize)
		n, err := reader.Read(chunk)
		if n > 0 {
			wg.Add(1)
			go func(order int, data []byte) {
				defer wg.Done()
				filteredChunk := filterWhitespace(data[:n])
				mu.Lock()
				chunkMap[order] = filteredChunk
				mu.Unlock()
			}(chunkOrder, chunk)
			chunkOrder++
		}
		if err != nil {
			break
		}
	}

	wg.Wait()

	var result []byte
	for i := 0; i < chunkOrder; i++ {
		result = append(result, chunkMap[i]...)
	}

	return result
}

func filterWhitespace(bytes []byte) []byte {
	result := make([]byte, 0, len(bytes))
	for _, b := range bytes {
		if !isWhitespaceCharacter(b) {
			result = append(result, b)
		}
	}
	return result
}

func getByteArrayHash(bytes []byte) uint32 {
	return murmur.MurmurHash2(normalizedBytes(bytes), 1)
}

func normalizedBytes(bytes []byte) []byte {
	var newArray []byte
	for _, b := range bytes {
		if !isWhitespaceCharacter(b) {
			newArray = append(newArray, b)
		}
	}
	return newArray
}

func isWhitespaceCharacter(b byte) bool {
	return b == 9 || b == 10 || b == 13 || b == 32
}
