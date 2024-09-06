package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

type ChunkedFile struct {
	Name   string
	Size   int64
	Chunks []Chunk
}

type Chunk struct {
	Number int
	Data   []byte
	Hash   string
}

func chunkFile(filename string, keepData bool) (*ChunkedFile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			os.Exit(1)
		}
	}()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	f := &ChunkedFile{
		Name:   stat.Name(),
		Size:   stat.Size(),
		Chunks: make([]Chunk, 0),
	}

	chunkSize := 4 * 1024 * 1024 // 4 MB
	chunk := make([]byte, chunkSize)

	i := 0
	for {
		_, err = file.Read(chunk)
		if err != nil && err == io.EOF {
			fmt.Println(err)
			break
		}

		if len(chunk) == chunkSize {
			var data []byte
			if keepData {
				data = chunk
			}
			f.Chunks = append(f.Chunks, Chunk{
				Number: i,
				Data:   data,
				Hash:   fmt.Sprintf("%x", sha256.Sum256(chunk)),
			})
			if err != nil {
				return nil, err
			}

			i = i + 1
			chunk = make([]byte, chunkSize)
		}

	}

	var data []byte
	if keepData {
		data = chunk
	}
	f.Chunks = append(f.Chunks, Chunk{
		Number: i,
		Data:   data,
		Hash:   fmt.Sprintf("%x", sha256.Sum256(chunk)),
	})

	return f, nil
}

func writeFile(f *ChunkedFile, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, chunk := range f.Chunks {
		_, err := file.Write(chunk.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	f, err := chunkFile("/home/bustanil/Downloads/20240831_094558.mp4", true)
	//f, err := chunkFile("/home/bustanil/wp-config.php")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("filename:", f.Name)
	fmt.Println("size:", f.Size)

	fmt.Println("chunks:", len(f.Chunks))
	for _, chunk := range f.Chunks {
		fmt.Println(chunk.Number, len(chunk.Data), chunk.Hash)
	}

	err = writeFile(f, "/home/bustanil/Downloads/20240831_094558_out.mp4")
	if err != nil {
		fmt.Println(err)
	}
}
