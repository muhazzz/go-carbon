package cache

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/lomik/go-carbon/points"
)

func BenchmarkDump(b *testing.B) {
	metrics := 1000000
	pointsCount := 5

	c := New()
	c.SetMaxSize(uint32(pointsCount*metrics + 1))

	baseTimestamp := time.Now().Unix()

	for i := 0; i < metrics; i++ {
		for j := 0; j < pointsCount; j++ {
			c.Add(points.OnePoint(
				fmt.Sprintf("carbon.localhost.cache.size%d", i),
				42.15*float64(j),
				baseTimestamp+int64(j),
			))
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		func() {
			tmpDir, err := ioutil.TempDir("", "")
			if err != nil {
				b.Fatal(err)
			}

			f, err := os.Create(path.Join(tmpDir, "dump.txt"))
			if err != nil {
				b.Fatal(err)
			}
			w := bufio.NewWriterSize(f, 1048576) // 1Mb

			c.Dump(w)

			w.Flush()
			f.Close()

			// s, err := os.Stat(path.Join(tmpDir, "dump.txt"))
			// fmt.Println(s.Size())

			if err := os.RemoveAll(tmpDir); err != nil {
				b.Fatal(err)
			}
		}()
	}
}

func BenchmarkDumpBinary(b *testing.B) {
	metrics := 1000000
	pointsCount := 5

	c := New()
	c.SetMaxSize(uint32(pointsCount*metrics + 1))

	baseTimestamp := time.Now().Unix()

	for i := 0; i < metrics; i++ {
		for j := 0; j < pointsCount; j++ {
			c.Add(points.OnePoint(
				fmt.Sprintf("carbon.localhost.cache.size%d", i),
				42.15*float64(j),
				baseTimestamp+int64(j),
			))
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		func() {
			tmpDir, err := ioutil.TempDir("", "")
			if err != nil {
				b.Fatal(err)
			}

			f, err := os.Create(path.Join(tmpDir, "dump.txt"))
			if err != nil {
				b.Fatal(err)
			}
			w := bufio.NewWriterSize(f, 1048576) // 1Mb

			c.DumpBinary(w)

			w.Flush()
			f.Close()

			// s, err := os.Stat(path.Join(tmpDir, "dump.txt"))
			// fmt.Println(s.Size())

			if err := os.RemoveAll(tmpDir); err != nil {
				b.Fatal(err)
			}
		}()
	}

}
