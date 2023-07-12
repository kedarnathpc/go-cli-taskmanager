package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

// create bucket task
// the values(key, values) will be stored as byte slice the the db
var taskBucket = []byte("tasks")

// makes connection with database
var db *bolt.DB

// structure of a task
type Task struct {
	Key   int
	Value string
}

// open the database for operations
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	// opens the db for read/write, if not exits creates one
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// creates a task
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {

		// take the bucket
		b := tx.Bucket(taskBucket)

		// get the next sequence
		id64, _ := b.NextSequence()
		id = int(id64)

		// assign the key
		key := itob(id)

		// save the key/value pair to the bucket
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// returns all tasks
func AllTasks() ([]Task, error) {
	var tasks []Task

	// open db for read
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		// iterate over the keys with cursor
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{

				// convert the byteslice to integer and assign
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})

}

// converts int to byte slice which is accepted by the bolt db
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// converts byte slice to int
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
