# Store Inventory System

## Problem Statement
Design and implement store inventory system for a multi-location retail chain that tracks inventory across various stores.

## Requirements
1. Store Management
    * Add Store: Ability to add a new store to the system with a unique identifier and details.
    * Remove store: Ability to remove a store from the system.
2. Stock Management
    * Adding Stock: Admin can add stock of a product to a specified store.
    * Removing Stock: Admin can remove or update the stock of a product at a store.
    * Query Stock: Users can check the availability of a product across all stores or a specific store.
3. Inventory Transfer Requests
    * Forward Request: If a product is not available at a queried store, the system should automatically check other stores for availability and optionally forward the request for transferring stock between stores.
    * Transfer Stock: Admin can manually initiate a stock transfer from one store to another.
4. The system should efficiently handle concurrent queries and updates from multiple users.
5. The inventory system should be robust enough to handle edge cases like transferring more stock than available.

## Example
```
ADD_STORE 001 Walmart
REMOVE_STORE 002
ADD_STOCK 001 1001 500
REMOVE_STOCK 001 1001 200
QUERY_STOCK 1001
TRANSFER_STOCK 001 003 1001 300

STORE_ADDED 001
STORE_REMOVED 002
STOCK_ADDED 001 1001 500
STOCK_REMOVED 001 1001 200
STOCK_AVAILABLE 001 1001 300
TRANSFER_INITIATED 001 003 1001 300
```
