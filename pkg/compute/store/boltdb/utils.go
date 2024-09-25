package boltdb

import (
	"encoding/binary"

	bolt "go.etcd.io/bbolt"
)

// strToBytes converts a string to a byte slice
func strToBytes(s string) []byte {
	return []byte(s)
}

// uint64ToBytes converts an uint64 to a byte slice
func uint64ToBytes(i uint64) []byte {
	buf := make([]byte, 8) //nolint:gomnd
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

// Helper function to safely access buckets and nested buckets
func bucket(tx *bolt.Tx, bucketNames ...string) *bolt.Bucket {
	var b *bolt.Bucket
	for _, name := range bucketNames {
		if b == nil {
			b = tx.Bucket([]byte(name))
		} else {
			b = b.Bucket([]byte(name))
		}
		if b == nil {
			return nil
		}
	}
	return b
}

// isEmptyBucket checks if a bucket is empty
func isEmptyBucket(b *bolt.Bucket) bool {
	k, _ := b.Cursor().First()
	return k == nil
}