package goissue27415_test

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	bolt "github.com/coreos/bbolt"
)

func BenchmarkBboltWrites(b *testing.B) {
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

	var buf [8]byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint64(buf[:], uint64(i))

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
