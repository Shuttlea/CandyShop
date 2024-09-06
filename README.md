# CandyShop
 For start server go to the src/server forlder and execute next commands:
```
go build
```
```
./server
```
For compile the client execute following command in the src/client folder:
```
go build
```
Flags for using the client:
 - `-m` - money that you put to the vending machine
 - `-c` - count of candies that you want to buy
 - `-k` - type of the candies\
\
Example:
```
./candy-client -m 40 -c 2 -k AA
```
Candies types and costs:
- `CA` - Cool Eskimo: 10 cents
- `AA` - Apricot Aardvark: 15 cents
- `NT` - Natural Tiger: 17 cents
- `DE` - Dazzling 	Elderberry: 21 cents
- `YR` - Yellow Rambutan: 23 cents


