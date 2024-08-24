package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open(".db", 0600, nil)
	// defer db.close()
	if err != nil {
		log.Fatal(err)
	}
	user := map[string]string{
		"name": "siddharth",
		"age":  "36",
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("users"))
		if err != nil {
			return err
		}
		for k, v := range user {
			if err := bucket.Put([]byte(k), []byte(v)); err != nil {
				return err
			}
		}
		return nil

	})
	fmt.Println("Works!!")
}
