# Bank Problem

## Problem Statement
Create a simple banking system that allows users to:
1. Credit
2. Debit
3. Check Balance
4. Exit

The amount entered by user will be in `xD yC` format. For example: `10D 20C`. It can also be just `xD` or `yC` for example: `10D` or `20C`. The amount shown should also be in the same format.

### Problem Extension
1. Support for multiple users
2. Support for multiple currency

## Design Pattern
Design pattern used in this solution:
* Factory Method Pattern - The `NewAccountAccess` function acts as a factory method. It encapsulates the creation of account objects and returns them as an interface type AccountAccess. This hides the creation details and allows for the introduction of new AccountAccess types without changing the client code.
