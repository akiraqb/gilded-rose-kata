package main

import "testing"


/* 
	TODO:

	- Analyse requirements 
	- write test for legacy gode to check if code works in the first place 100%
	- thing of refactoring
	- add tests for refactor to aligh with previous logic, keep as much coverage as possible"
	- implement refactoring */


/* 
	Requirements (acording to discription)
	
	- there are 2 group so fitems:
		* normal items (everthing which is not special case):
			- both quality and sellin dicrease by 1 every day
			- quality item is necer negative 
			- quality item is never more than 50
			- once sellin date pass quality degraeds twice as fast

		* special 

*/





func Test_GildedRose(t *testing.T) {
	

	var items = []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}



	t.Run("Test_main", func(t *testing.T){ main() })
	t.Run("First_test",func(t *testing.T){ GildedRose(items) })




}
