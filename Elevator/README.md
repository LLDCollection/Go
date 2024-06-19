# Elevator
Low level design and implementation for Elevator

## Problem Statement
Design a system for managing the operation of an elevator in a multi-story building.

## Requirements
1.	Basic Operation:
	* The elevator should be able to move up and down between floors.
	* The elevator should stop at requested floors to pick up and drop off passengers.
2.	Floor Requests:
    * Passengers can request the elevator from any floor by pressing up or down buttons.
    * Inside the elevator, passengers can select their desired floor.
3.	Direction and Movement:
    * The elevator should continue in its current direction until there are no more requests in that direction before reversing.
    * The elevator should prioritize stops based on the current direction of travel to optimize efficiency.
4.	Capacity Management:
    * The elevator should have a maximum weight capacity. It should not accept more passengers if the capacity is reached.
    * Display the current load and capacity status.

## Extensions
1.	Multiple Elevators: If there are multiple elevators in the building, optimize the allocation of elevator requests to minimize wait time and energy consumption.

## Appendix
### Clarification Questions
1.	How many floors does the building have?
2.	Is there a specific time frame in which the elevator should be available (e.g., 24/7 service)?
3.	Are there any peak hours where elevator usage is significantly higher, requiring optimization?
4.	How many elevators are there in the building, and should they work together as a system?
5.	What is the maximum weight capacity of the elevator?
6.	Are there any special requirements for accessibility (e.g., wheelchair access)?

### Design Suggestions
* Use a priority queue or similar data structure to manage and prioritize floor requests.
* Implement a state machine to handle the elevator’s current state (e.g., moving up, moving down, idle).
* Consider thread safety and synchronization mechanisms to manage concurrent requests and avoid race conditions.