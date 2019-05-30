package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
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


func Test_GildedRose(t *testing.T){

	t.Run("Update_sellIn_quality_value_add", func(t *testing.T) {
		//create array of testdata

		items := []Item{{"+5 Dexterity Vest", 10, 10}}
		items[0].updateSellIn(+1)
		items[0].updateQuality(+1)
		//expected actual
		assert.Equal(t, 11, items[0].sellIn, "wrong sellin value")
		assert.Equal(t, 11, items[0].quality, "wrong sellin value")

	})
	t.Run("Update_sellIn_quality_value_substract", func(t *testing.T) {
		//create array of testdata

		items := []Item{{"+5 Dexterity Vest", 10, 10}}
		items[0].updateSellIn(-1)
		items[0].updateQuality(-1)
		//expected actual
		assert.Equal(t, 9, items[0].sellIn, "wrong sellin value")
		assert.Equal(t, 9, items[0].quality, "wrong sellin value")

	})

}


func Test_Normal_Update(t *testing.T){

	t.Run("Normal_Update_sellIn_decrease", func(t *testing.T){

			Nom := CreateNormal(&Item{"+5 Dexterity Vest", 10, 10})
			Nom.Update()
			assert.Equal(t,9,Nom.sellIn,"error sellin vaule")
		})

	t.Run("Normal_Update_quality_decrease", func(t *testing.T){

			Nom := CreateNormal(&Item{"+5 Dexterity Vest", 10, 10})
			Nom.Update()
			assert.Equal(t,9,Nom.quality,"error sellin vaule")
		})

	t.Run("Normal_Update_quality_never_negative", func(t *testing.T){

			Nom := CreateNormal(&Item{"+5 Dexterity Vest", 10, 0})
			Nom.Update()
			assert.Equal(t,9,Nom.sellIn,"error sellin vaule")
			assert.Equal(t,0,Nom.quality,"error quality value ")
		})

	t.Run("Normal_Update_quality_decrease_twice", func(t *testing.T){

			Nom := CreateNormal(&Item{"+5 Dexterity Vest", -1, 10})
			Nom.Update()
			assert.Equal(t,-2,Nom.sellIn,"error sellin vaule")
			assert.Equal(t,8,Nom.quality,"error quality value ")
		})

	t.Run("Normal_Update_quality_max", func(t *testing.T){

			Nom := CreateNormal(&Item{"+5 Dexterity Vest", 10, 55})
			Nom.Update()
			assert.Equal(t,9,Nom.sellIn,"error sellin vaule")
			assert.Equal(t,50,Nom.quality,"error quality value ")
		})

} // end of normal
func Test_Backstage_Update(t *testing.T){

			// Backstage
		t.Run("Backstage_Update_sellIn_decrease", func(t *testing.T){

			BS:=CreateBackstage(&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 10})
			BS.Update()
			assert.Equal(t,14,BS.sellIn,"error sellin vaule")
		})


		t.Run("Backstage_Update_quality_increase_by_2", func(t *testing.T){

			BS:=CreateBackstage(&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 10})
			BS.Update()
			assert.Equal(t,9,BS.sellIn,"error sellin vaule")
			assert.Equal(t,12,BS.quality,"error quality value ")
		})

		t.Run("Backstage_Update_quality_increase_by_3", func(t *testing.T){

			BS:=CreateBackstage(&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 10})
			BS.Update()
			assert.Equal(t,4,BS.sellIn,"error sellin vaule")
			assert.Equal(t,13,BS.quality,"error quality value ")
		})

		t.Run("Backstage_Update_quality_drops_zero", func(t *testing.T){

			BS:=CreateBackstage(&Item{"Backstage passes to a TAFKAL80ETC concert", 0, 10})
			BS.Update()
			assert.Equal(t,-1,BS.sellIn,"error sellin vaule")
			assert.Equal(t,0,BS.quality,"error quality value ")
		})

		t.Run("Backstage_Update_quality_max", func(t *testing.T){

			BS:=CreateBackstage(&Item{"Backstage passes to a TAFKAL80ETC concert", 25, 15})
			//shouuld incrase up to max 50
			for i:= 0; i < 24 ; i++ {
				BS.Update()
			}	
			assert.Equal(t,1,BS.sellIn,"error sellin vaule")
			assert.Equal(t,50,BS.quality,"error quality value ")
		})


}

func Test_AgedBrie_Update(t *testing.T){

	
		t.Run("AgedBrie_Update_sellIn_decrease", func(t *testing.T){

			ABrie:= CreateAgedBrie(&Item{"Aged Brie", 10, 10})
			ABrie.Update()
			assert.Equal(t,9,ABrie.sellIn,"error sellin vaule")
		})

		t.Run("AgedBrie_Update_quality_increase", func(t *testing.T){

			ABrie:= CreateAgedBrie(&Item{"Aged Brie", 10, 10})
			ABrie.Update()
			assert.Equal(t,11,ABrie.quality,"error quality value ")
		})

		t.Run("AgedBrie_Update_quality_max", func(t *testing.T){

			ABrie:= CreateAgedBrie(&Item{"Aged Brie", 60, 10})
			//shouuld incrase up to max 50
			for i:= 0; i < 50 ; i++ {
				ABrie.Update()
			}
			assert.Equal(t,10,ABrie.sellIn,"error sellin vaule")
			assert.Equal(t,50,ABrie.quality,"error quality value ")
		})

}

func Test_LegacyGildedRose(t *testing.T) {

	// var items = []Item{
	// 	Item{"+5 Dexterity Vest", 10, 20},
	// 	Item{"Aged Brie", 2, 0},
	// 	Item{"Elixir of the Mongoose", 5, 7},
	// 	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	// 	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	// 	Item{"Conjured Mana Cake", 3, 6},
	// }

	t.Run("Test_main", func(t *testing.T) { main() })
	

	/* normal items */

	t.Run("Normal_sellIn_decrease", func(t *testing.T) {
		//create array of testdata

		items := []Item{{"+5 Dexterity Vest", 10, 10}}
		GildedRose(items)
		//expected actual
		assert.Equal(t, 9, items[0].sellIn, "wrong sellin value")

	})
	t.Run("Normal_quality_decrease", func(t *testing.T) {

		items := []Item{{"+5 Dexterity Vest", 10, 10}}
		GildedRose(items)
		//expected actual
		assert.Equal(t, 9, items[0].quality, "wrong quality value")
	})
	t.Run("Normal_quality_never_negative", func(t *testing.T) {

		items := []Item{{"+5 Dexterity Vest", 10, 0}}
		GildedRose(items)
		//expected actual
		assert.Equal(t, 0, items[0].quality, "wrong quality value")
	})

	/* WARNING!  Quality of normal item never rises 
		so it will not apply, Unless initial value > 50  */

	// t.Run("Normal_quality_max", func(t *testing.T){

	// 	items:=[]Item{{"+5 Dexterity Vest", 10, 53},}
	// 	GildedRose(items)
	// 				//expected actual
	// 	assert.Equal(t,50, items[0].quality,"wrong quality value")
	// })

	t.Run("Normal_quality_decrease_twice_as_fast", func(t *testing.T) {

		items := []Item{{"+5 Dexterity Vest", -1, 10}}
		GildedRose(items)
		//expected actual
		assert.Equal(t, 8, items[0].quality, "wrong quality value")
	})

	/* Special */
	/* Aged Bri items */

	t.Run("Aged_Brie_sellIn_decrease", func(t *testing.T) {

		items := []Item{{"Aged Brie", 10, 10}}
		GildedRose(items)
		//expected actual
		assert.Equal(t, 9, items[0].sellIn, "wrong quality value")
	})

	t.Run("Aged_Brie_quality_increase", func(t *testing.T) {

		items := []Item{{"Aged Brie", 10, 10}}
		GildedRose(items)
		//expected actual
		assert.Equal(t, 11, items[0].quality, "wrong quality value")
	})

	t.Run("Aged_Brie_quality_max", func(t *testing.T) {

		items := []Item{{"Aged Brie", 60, 45}}
		for i := 0; i < 20; i++ {
			GildedRose(items)
		}
		//expected actual
		assert.Equal(t, 40, items[0].sellIn, "wrong quality value")
		assert.Equal(t, 50, items[0].quality, "error quality value ")
	})

	/* WARNING
		Requirements say it should *decrease* for normla item
		Code suggest its continues to increase?	*/

	t.Run("Aged_Brie_quality_increses_twice_as_fast", func(t *testing.T){

		items:=[]Item{{"Aged Brie", -1, 10},}
		GildedRose(items)
					//expected actual
		assert.Equal(t,12, items[0].quality,"wrong quality value")

		// i think it should actually decrease
		// assert.Equal(t,8, items[0].quality,"wrong quality value")

	})

	/* Backstage passes */

	t.Run("Backstage_sellIn_decrease", func(t *testing.T) {

		items := []Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 10}}
		GildedRose(items)
		assert.Equal(t, 9, items[0].sellIn, "wrong quality value")
	})

	t.Run("Backstage_quality_increase_by_2", func(t *testing.T) {

		items := []Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 10}}
		GildedRose(items)
		assert.Equal(t, 9, items[0].sellIn, "wrong quality value")
		assert.Equal(t, 12, items[0].quality, "wrong quality value")
	})

	t.Run("Backstage_quality_increase_by_3", func(t *testing.T) {

		items := []Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 10}}
		GildedRose(items)
		assert.Equal(t, 4, items[0].sellIn, "wrong quality value")
		assert.Equal(t, 13, items[0].quality, "wrong quality value")
	})

	t.Run("Backstage_quality_decrease_to_0", func(t *testing.T) {

		items := []Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 10}}
		GildedRose(items)
		assert.Equal(t, -1, items[0].sellIn, "wrong quality value")
		assert.Equal(t, 0, items[0].quality, "wrong quality value")
	})

	t.Run("Backstage_quality_,ax", func(t *testing.T) {

		items := []Item{{"Backstage passes to a TAFKAL80ETC concert", 25, 15}}

		for i := 0; i < 24; i++ {
			GildedRose(items)
		}
		assert.Equal(t, 1, items[0].sellIn, "wrong quality value")
		assert.Equal(t, 50, items[0].quality, "wrong quality value")
	})

	/* Sulfuras*/

	t.Run("Sulfuras_sellIn_never_decrease", func(t *testing.T) {

		items := []Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}

		GildedRose(items)
		assert.Equal(t, 0, items[0].sellIn, "error sellin vaule")
	})

	t.Run("Sulfuras_quality_never_decrease", func(t *testing.T) {
		items := []Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}
		GildedRose(items)
		assert.Equal(t, 80, items[0].quality, "error quality value ")
	})

}
