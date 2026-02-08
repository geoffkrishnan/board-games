package main

/*
* Game loop:
* 	create two boards & two sets of ships
* 	pick a player at random to go first
* 	setup phase for both boards
* 		within 30 seconds, both players concurrently place their ships on their boards
* 		so i think we can create two arrays of the ships we want players to place
* 		while that array is still full and 30 seconds have not passed:
* 			user can pick a ship from that array.
* 			they can drag it around the grid
* 			they can rotate it as well
* 			they can put it back into the array or place it on the grid.
* 			we could just have it lock into place once its on the grid and remove it from the array.
* 			OR we could make it so all the ships will lock into place once and are removed from the array once the 30 seconds are up
* 		when the 30 seconds have passed AND while the array is not empty
* 			pick a random ship and place it on a random valid space on grid and remove it from the array
*
* 	turn by turn game phase until game over
* 		each turn is 15 seconds(eventually make it configurable in room/game settings by players probably but just hardcode it for the mvp)
* 		player 1 - chooses a valid square on enemy board
* 			a square is valid if 
* 				hasn't been chosen before. that's pretty much the only check we need to do
* 			so have to keep track of if a square has been visited. visited set, this pattern shows up a lot
* 			check if that square is a hit on the enemy board
* 			if it is mark that hit on the coords of ship object that was hit
* 			otherwise just mark it as visited so we can't choose that square again
* 			check if player 1 has won, if not, swap player and loop until either player wins
*/

func main() {
	player1Board := GetBoard()
	player2Board := GetBoard()

	shipSizes := []int{5, 4, 3, 3, 2}
	player1Ships := []Ship{}
	player2Ships := []Ship{}


	for _, size := range shipSizes {
		player1Ships = append(player1Ships, Ship{
			Size: size,
			Coords: make(map[Coord]Hit),
			Sunk: false,
		})
		player2Ships = append(player2Ships, Ship{
			Size: size,
			Coords: make(map[Coord]Hit),
			Sunk: false,
		})
	}
}
