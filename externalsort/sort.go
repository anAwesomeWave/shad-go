//go:build !solution

package externalsort

import (
	"container/heap"
	"io"
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
		if byteCnt == 0 || readErr == io.EOF || charSlice[0] == '\n' {
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

type StrHeap []string

func (h StrHeap) Len() int           { return len(h) }
func (h StrHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h StrHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StrHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(string))
}

func (h *StrHeap) Pop() any {
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
	//allStrs := make([]string, 0, len(readers)) // cap min -> len(readers)
	strHeap := StrHeap{}
	for _, r := range readers {
		for { // eof -> end
			str, rErr := r.ReadLine()
			//allStrs = append(allStrs, str)
			heap.Push(&strHeap, str)
			if rErr != nil {
				break
			}
		}
	}
	for len(strHeap) > 0 {
		elem := heap.Pop(&strHeap).(string)
		if err := w.Write(elem); err != nil {
			return err
		}
	}
	return nil
}

func Sort(w io.Writer, in ...string) error {
	/*
		получаем файлы. их превращаем в line readers сортируем между собой и отдаем в merge
	*/
	sort.Strings(in)
	for _, str := range in {
		_, err := w.Write([]byte(str))
		if err != nil {
			return err
		}
	}
	return nil
}
