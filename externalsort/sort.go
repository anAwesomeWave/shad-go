//go:build !solution

package externalsort

import (
	"container/heap"
	"io"
	"os"
	"sort"
)

type MyLReader struct {
	r io.Reader
}

func (lineR MyLReader) ReadLine() (string, error) {
	data := make([]byte, 0) // yes. 100% bad without cap
	charSlice := make([]byte, 1)
	var readErr error
	var byteCnt int
	for {
		byteCnt, readErr = lineR.r.Read(charSlice)
		if byteCnt != 0 && charSlice[0] != '\n' {
			data = append(data, charSlice...)
		}
		if byteCnt == 0 || readErr != nil || charSlice[0] == '\n' {
			break
		}
	}
	return string(data), readErr
}

type MyLWriter struct {
	w io.Writer
}

func (lineW MyLWriter) Write(p string) error {
	pLine := append([]byte(p), '\n')
	_, err := lineW.w.Write(pLine)
	return err
}

func NewReader(r io.Reader) LineReader {
	return MyLReader{r}
}

func NewWriter(w io.Writer) LineWriter {
	return MyLWriter{w: w}
}

type DataReader struct {
	line   string
	reader LineReader
}

type Heap []*DataReader

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].line < h[j].line }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*DataReader))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Merge(w LineWriter, readers ...LineReader) error {
	/*
		каждый line reader посорчен. этап сортировки всех посл. в 1
	*/
	pq := Heap{}
	for _, r := range readers {
		line, err := r.ReadLine()
		if err == io.EOF && len(line) == 0 {
			continue
		}
		heap.Push(&pq, &DataReader{line: line, reader: r})

	}
	for len(pq) > 0 {
		elem := heap.Pop(&pq).(*DataReader)

		if err := w.Write(elem.line); err != nil {
			return err
		}
		newLine, err := elem.reader.ReadLine()

		if err != nil && err != io.EOF {
			return err
		}

		if err == io.EOF && len(newLine) == 0 {
			continue
		}
		elem.line = newLine
		heap.Push(&pq, elem)
	}
	return nil
}

func Sort(w io.Writer, in ...string) error {
	/*
		получаем файлы. их превращаем в line readers сортируем между собой и отдаем в merge
	*/
	// 1. get file
	// 2. sort
	// 3. remove from memo
	for _, filePath := range in {

		rfile, err := os.OpenFile(filePath, os.O_RDONLY, 7555)

		if err != nil {
			return err
		}
		defer rfile.Close()

		lReader := NewReader(rfile)
		data := []string{}
		for {
			newStr, err := lReader.ReadLine()
			if err == io.EOF && len(newStr) == 0 {
				break
			}
			data = append(data, newStr)
		}
		if err := rfile.Close(); err != nil {
			return err
		}
		if len(data) == 0 {
			continue
		}
		wfile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		defer wfile.Close()
		sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
		lfWriter := NewWriter(wfile)
		for _, str := range data {
			if err := lfWriter.Write(str); err != nil {
				return err
			}
		}
		wfile.Close()
	}
	// 4. merge all
	readers := make([]LineReader, 0, len(in))
	for _, filePath := range in {
		file, err := os.OpenFile(filePath, os.O_RDONLY, 'r')
		defer file.Close()
		if err != nil {
			return err
		}
		readers = append(readers, NewReader(file))
	}
	lWriter := NewWriter(w)
	return Merge(lWriter, readers...)

}
