package store

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreImpl_AddStock(t *testing.T) {
	t.Run("when stock is added it should be present in inventory", func(t *testing.T) {
		store := NewStore("store1")

		store.AddStock([]Product{{Id: "1", Name: "A"}})

		q := store.QueryStock("A")
		assert.Equal(t, StockResult{StoreId: "store1", Count: 1}, q)
	})

	t.Run("when stock is added concurrently, it should correctly present in inventory", func(t *testing.T) {
		store := NewStore("store1")

		var wg sync.WaitGroup
		wg.Add(100)

		for i := 0; i < 100; i++ {
			go func(idx int) {
				defer wg.Done()
				store.AddStock([]Product{{Id: fmt.Sprintf("%v", idx), Name: "A"}})
			}(i)
		}

		wg.Wait()

		q := store.QueryStock("A")
		assert.Equal(t, StockResult{StoreId: "store1", Count: 100}, q)
	})
}
