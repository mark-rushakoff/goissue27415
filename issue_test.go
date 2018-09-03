package goissue27415_test

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"testing"

	bolt "github.com/coreos/bbolt"
)

func BenchmarkBboltWrites(b *testing.B) {
	parallelCount := []int{10, 100, 1000}
	for _, c := range parallelCount {
		b.Run(fmt.Sprint(c), func(b *testing.B) {
			dir, err := ioutil.TempDir("", "")
			if err != nil {
				b.Fatal(err)
			}

			path := filepath.Join(dir, "bolt.db")
			defer os.Remove(path)

			db, err := bolt.Open(path, 0644, nil)
			if err != nil {
				b.Fatal(err)
			}
			defer db.Close()

			ch := make(chan uint64, c)
			var wg sync.WaitGroup
			wg.Add(c)
			for i := 0; i < c; i++ {
				go write(b, db, ch, &wg)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ch <- uint64(i)
			}
			close(ch)
			wg.Wait()
		})
	}
}

func write(b *testing.B, db *bolt.DB, ch <-chan uint64, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range ch {
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], n)

		if err := db.Update(func(tx *bolt.Tx) error {
			// Create a bucket.
			bucket, err := tx.CreateBucketIfNotExists([]byte("numbers"))
			if err != nil {
				return err
			}

			if err := bucket.Put(buf[:], []byte("ok")); err != nil {
				return err
			}
			return nil
		}); err != nil {
			b.Error(err)
		}
	}
}
