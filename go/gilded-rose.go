package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

// var items = []Item{
// 	Item{"+5 Dexterity Vest", 10, 20},
// 	Item{"Aged Brie", 2, 0},
// 	Item{"Elixir of the Mongoose", 5, 7},
// 	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
// 	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
// 	Item{"Conjured Mana Cake", 3, 6},
// }


/* 
	Design:
	- change items object of its class.
	- ensure all can be updated
	- loop through them using polymorphic processing (pdate them all)
	- new test/new code agains the same desired logic 
*/

/* Functions for basic increase/decrease operations */

						//+= -1  will substrack
						//+= +1  will add
func (i * Item) updateSellIn(value int){
	i.sellIn += value
}

func (i * Item) updateQuality(value int){
	i.quality += value
}

/* Add subtypes of objects */

type Normal struct{
	*Item
}
type AgedBrie struct{
	*Item
}
type Sulfuras struct{
	*Item
}
type Backstage struct{
	*Item
}

const maxQuality = 50
const minQuality = 0

/* constructors */

func CreateNormal(item *Item) *Normal{
	//pointer to our type
	pNormal:= &Normal{Item:item,}
	return pNormal
}

func CreateBackstage(item *Item) *Backstage{

	pBs:= &Backstage{Item:item,}
	return pBs
}

/* Update logic*/
func (item *Normal) Update(){

	item.updateSellIn(-1)
	if item.sellIn < 0 {
		item.updateQuality(-2)
	}else{
		item.updateQuality(-1)
	}
	if item.quality < minQuality{
		item.quality = minQuality
	}
	
	if item.quality > maxQuality {
		item.quality = maxQuality
	}
} //end of Normal.update()


func (item *Backstage) Update(){
	item.updateSellIn(-1)
	sellIn:=item.sellIn

	switch{
	case sellIn >=10:
		item.updateQuality(+1)
	case sellIn <10 && sellIn > 5:
		item.updateQuality(+2)
	case sellIn <=5 && sellIn > 0:
		item.updateQuality(+3)
	case sellIn < 0:
		item.quality = minQuality
	}
	if item.quality > maxQuality{
		item.quality = maxQuality
	}


}






func main() {
	fmt.Println("OMGHAI!")
	// fmt.Print(items)
	// GildedRose()
}

func GildedRose(items []Item) {
	for i := 0; i < len(items); i++ {

		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].quality > 0 {
				if items[i].name != "Sulfuras, Hand of Ragnaros" {
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != "Aged Brie" {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].quality > 0 {
						if items[i].name != "Sulfuras, Hand of Ragnaros" {
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}

}
