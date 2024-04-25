# Parking Lot
Low level design for a Parking Lot.

## Problem Statement
Design a parking lot.

### Requirements
1. Different Spot Sizes: There are different sizes of parking spots to accommodate different types of vehicles.
    a. Sizes: S, M, L, XL
    b. A smaller size vehicle can be parked in a larger spot but not the vice versa.
2. Automated Ticketing: On entering, vehicles should be issued a ticket that records the time of entry, and charges should be calculated at the exit based on the duration of the stay.
3. Parking Spot Assignment: The system should automatically find the nearest available parking spot to minimize the distance the driver needs to walk.
4. Occupancy Display: Real-time display of available spots on each floor or section.
5. Multiple Floors: The parking lot has multiple floors with each floor having various sections.

#### Extension
1. Multiple Entry and Exit.

## Appendix
### Clarification Questions
#### Business Requirement clarification
1. How many parking spots are we looking at? 10s? 1000s?
2. Is this parking lot a building? Are there multiple levels?
3. Do we fill in certain parking spots before others?
3. Are there multiple entrances? (Concurrency issues)
4. Are all parking spots same or they are of different sizes?
5. Do we want to park closer to the entrance or the exit?
6. Do we charge for the parking? Does the charge increase with time?