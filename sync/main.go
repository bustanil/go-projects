package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

type File struct {
	Name   string
	Size   int64
	Chunks []Chunk
}

type Chunk struct {
	Number int
	Data   []byte
}

func chunkFile(filename string) (*File, error) {
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

	f := &File{
		Name:   stat.Name(),
		Size:   stat.Size(),
		Chunks: make([]Chunk, 0),
	}

	chunkSize := 4 * 1024 * 1024 // 4 MB
	chunk := make([]byte, 0)

	bufferSize := 2 * 1024 * 1024
	buffer := make([]byte, bufferSize)

	i := 0
	var bytesRead int
	for {
		bytesRead, err = file.Read(buffer)
		chunk = slices.Concat(chunk, buffer[:bytesRead])
		if err != nil && err == io.EOF {
			fmt.Println(err)
			break
		}

		if len(chunk) == chunkSize {
			f.Chunks = append(f.Chunks, Chunk{
				Number: i,
				Data:   chunk,
			})
			if err != nil {
				return nil, err
			}

			i = i + 1
			bytesRead = 0
			chunk = make([]byte, 0)
		}

	}

	f.Chunks = append(f.Chunks, Chunk{
		Number: i,
		Data:   chunk,
	})

	return f, nil
}

func main() {
	f, err := chunkFile("/home/bustanil/Downloads/20240831_094558.mp4")
	//f, err := chunkFile("/home/bustanil/wp-config.php")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("filename:", f.Name)
	fmt.Println("size:", f.Size)

	fmt.Println("chunks:", len(f.Chunks))
	for _, chunk := range f.Chunks {
		fmt.Println(chunk.Number, len(chunk.Data))
	}
}
