package exercise

import "fmt"

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

// pick up an item and add it to the inventory
func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s\n", p.Name, item.Name)
}

// drop an item and remove it from the inventory
func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s\n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("%s does not have %s in their inventory\n", p.Name, itemName)
}

//Use an item ( if potion , remove it after use)
func (p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
				fmt.Printf("%s used %s\n", p.Name, itemName)
				return
			}
		}
	}
	fmt.Printf("%s does not have %s in their inventory\n", p.Name, itemName)
}

// list all items in the inventory
func (p *Player) ListItems() {
	for _, item := range p.Inventory {
		fmt.Printf("%s: %s\n", item.Name, item.Type)
	}
}

// Exercise function to test the inventory system
func Exercise() {
	player := Player{Name: "John"}
	player.PickUpItem(Item{Name: "Sword", Type: "weapon"})
	player.PickUpItem(Item{Name: "Shield", Type: "armor"})
	player.PickUpItem(Item{Name: "Potion", Type: "potion"})
	player.ListItems()
	player.UseItem("Potion")
	player.DropItem("Shield")
	player.ListItems()
}
