package store

import (
	"errors"
	"fmt"
	"sync"
)

type Store interface {
	AddStock(products []Product)
	RemoveStock(name string, number int)
	QueryStock(product string) StockResult

	SetNextStore(store Store)
}

type StoreImpl struct {
	Id        string
	Address   Address
	Inventory map[string][]Product
	Next      Store

	Lock sync.RWMutex
}

type Product struct {
	Id   string
	Name string
}

type InventoryItem struct {
	Id string
	Product
}

type Address struct {
	Street  string
	City    string
	State   string
	Zipcode string
}

func (s *StoreImpl) AddStock(products []Product) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	for _, prd := range products {
		s.Inventory[prd.Name] = append(s.Inventory[prd.Name], prd)
	}
}

func (s *StoreImpl) RemoveStock(productName string, number int) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	stock, ok := s.Inventory[productName]
	if !ok {
		return errors.New("product not available")
	}

	if len(stock) < number {
		return fmt.Errorf("product count is only %v but removal request sent for %v", len(stock), number)
	}

	s.Inventory[productName] = s.Inventory[productName][:len(stock)-number]

	return nil
}

func (s *StoreImpl) QueryStock(productName string) StockResult {
	s.Lock.RLock()
	stock, ok := s.Inventory[productName]
	s.Lock.RUnlock()

	if ok {
		return StockResult{StoreId: s.Id, Count: len(stock)}
	}

	if s.Next != nil {
		return s.Next.QueryStock(productName)
	}

	return StockResult{}
}

func (s *StoreImpl) SetNextStore(store Store) {
	s.Next = store
}
