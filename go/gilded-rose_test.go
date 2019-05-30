package main

import "testing"

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
