package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


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
			- "Aged Brie" increases in quality the older it gets
			- "Backstage passes" increase in qualiuty as it sellIn older it gets

				- quality increases by 2 when there are 10 days or less
				- quality increases by 3 when there are  5 days or less
				- quality drops to 0  afterdate 

			- "Sulfuras" being a legendary item, never has to be sold or decreased in quality,
				fixed value of 80  
			- "Conjured" items degrade in quality twice as fast as normal items.
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

	/* normal items */


	t.Run("Normal_sellIn_decrease", func(t *testing.T){ 
		//create array of testdata

		items:=[]Item{{"+5 Dexterity Vest", 10, 10},}
		GildedRose(items)
					//expected actual
		assert.Equal(t,9, items[0].sellIn,"wrong sellin value")

	})
	t.Run("Normal_quality_decrease", func(t *testing.T){ 
		
		items:=[]Item{{"+5 Dexterity Vest", 10, 10},}
		GildedRose(items)
					//expected actual
		assert.Equal(t,9, items[0].quality,"wrong quality value")
	})
	t.Run("Normal_quality_never_negative", func(t *testing.T){ 

		items:=[]Item{{"+5 Dexterity Vest", 10, 0},}
		GildedRose(items)
					//expected actual
		assert.Equal(t,0, items[0].quality,"wrong quality value")
	})

	/* Quality of normal item never rises so it will not apply
		Its a staring value in the system */
	// t.Run("Normal_quality_never_more_50", func(t *testing.T){ 

	// 	items:=[]Item{{"+5 Dexterity Vest", 10, 53},}
	// 	GildedRose(items)
	// 				//expected actual
	// 	assert.Equal(t,50, items[0].quality,"wrong quality value")
	// })

	t.Run("Normal_quality_decrease_twice_as_fast", func(t *testing.T){ 

		items:=[]Item{{"+5 Dexterity Vest", -1, 10},}
		GildedRose(items)
					//expected actual
		assert.Equal(t,8, items[0].quality,"wrong quality value")
	})

	/* normal items */







}
