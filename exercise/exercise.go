package main

import (
	"fmt"
	"math/rand"
)

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name string
	Inventory []Item
}

func NewInventories() []Item {
	defaultItem := Items[len(Items) - 1]
	return []Item{defaultItem}
}

 var Items = []Item {
	{Name: "Weapons1", Type: "Tes1"},
	{Name: "Weapons2", Type: "Tes2"},
	{Name: "Weapons3", Type: "Tes3"},
	{Name: "Weapons4", Type: "Tes4"},
	{Name: "Weapons5", Type: "Tes5"},
	{Name: "Weapons6", Type: "Tes6"},
	{Name: "Weapons7", Type: "Tes7"},
	{Name: "Weapons8", Type: "Tes7"},
	{Name: "Weapons9", Type: "Tes8"},
	{Name: "Weapons10", Type: "Te1"},
	{Name: "Weapons11", Type: "Tes1"},
	{Name: "Weapons12", Type: "Tes3"},
	{Name: "Weapons13", Type: "Tes4"},
	{Name: "Weapons14", Type: "Tes5"},
	{Name: "Weapons15", Type: "Tes6"},
	{Name: "Weapons16", Type: "Tes7"},
	{Name: "Weapons17", Type: "Tes8"},
	{Name: "Weapons18", Type: "Tes9"},
}

func NewPlayer() Player {
	return Player {
		Name: "Dmitrii",
		Inventory: NewInventories(),
	}
}

func (player *Player) PickUpItem(item Item) {
	player.Inventory = append(player.Inventory, item)
}

func (player *Player) DropItem(itemName string) []Item {
	result := []Item{}
	for _, item := range player.Inventory {
		if item.Name != itemName {
			result = append(result, item)
		}
	}

	player.Inventory = result

	return result
}

func (player *Player) UseItem(itemName string) {
	for _, item := range player.Inventory {
		if item.Name == itemName {
			fmt.Printf("Print Random Inventory Name: %s and Kind %s \n", item.Name, item.Type)
			return;
		}
	}

	fmt.Printf("No Found")
}

func main() {
	player :=  NewPlayer();
	player.PickUpItem(Items[rand.Intn(len(Items))])
	player.PickUpItem(Items[rand.Intn(len(Items))])
	player.PickUpItem(Items[rand.Intn(len(Items))])
	player.PickUpItem(Items[rand.Intn(len(Items))])

	fmt.Printf("New Player %v \n", player)
	deleteRandIdx := rand.Intn(len(player.Inventory));
	player.DropItem(player.Inventory[deleteRandIdx].Name)
	fmt.Printf("New Player %v \n", player)


	useRandIdx := rand.Intn(len(player.Inventory));
	player.UseItem(player.Inventory[useRandIdx].Name)

}