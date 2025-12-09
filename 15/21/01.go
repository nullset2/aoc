package main

import (
	"fmt"
	"math"
)

type Player struct {
	HP  int
	Dmg int
	Def int
}

type Item struct {
	Name string
	Cost int
	Dmg  int
	Def  int
}

type Loadout struct {
	Items     []Item
	TotalCost int
	TotalDmg  int
	TotalDef  int
}

func generateCombinations(itemShop, armorShop, ringShop []Item) []Loadout {
	var combinations []Loadout

	// Helper function to create a loadout from selected items
	createLoadout := func(items ...Item) Loadout {
		loadout := Loadout{Items: items}
		for _, item := range items {
			loadout.TotalCost += item.Cost
			loadout.TotalDmg += item.Dmg
			loadout.TotalDef += item.Def
		}
		return loadout
	}

	// Rules:
	// - Exactly 1 weapon (required)
	// - 0 or 1 armor (optional)
	// - 0, 1, or 2 rings (optional, can't buy same ring twice)

	for _, weapon := range itemShop {
		// 1. Weapon only (no armor, no rings)
		combinations = append(combinations, createLoadout(weapon))

		// 2. Weapon + Armor (no rings)
		for _, armor := range armorShop {
			combinations = append(combinations, createLoadout(weapon, armor))
		}

		// 3. Weapon + 1 Ring (no armor)
		for _, ring := range ringShop {
			combinations = append(combinations, createLoadout(weapon, ring))
		}

		// 4. Weapon + 2 Rings (no armor)
		for i, ring1 := range ringShop {
			for j, ring2 := range ringShop {
				if i < j { // Ensure we don't buy the same ring twice
					combinations = append(combinations, createLoadout(weapon, ring1, ring2))
				}
			}
		}

		// 5. Weapon + Armor + 1 Ring
		for _, armor := range armorShop {
			for _, ring := range ringShop {
				combinations = append(combinations, createLoadout(weapon, armor, ring))
			}
		}

		// 6. Weapon + Armor + 2 Rings
		for _, armor := range armorShop {
			for i, ring1 := range ringShop {
				for j, ring2 := range ringShop {
					if i < j { // Ensure we don't buy the same ring twice
						combinations = append(combinations, createLoadout(weapon, armor, ring1, ring2))
					}
				}
			}
		}
	}

	return combinations
}

func main() {
	itemShop := []Item{
		Item{Name: "Dagger", Cost: 8, Dmg: 4, Def: 0},
		Item{Name: "Shortsword", Cost: 10, Dmg: 5, Def: 0},
		Item{Name: "Warhammer", Cost: 25, Dmg: 6, Def: 0},
		Item{Name: "Longsword", Cost: 40, Dmg: 7, Def: 0},
		Item{Name: "Greataxe", Cost: 74, Dmg: 8, Def: 0},
	}

	armorShop := []Item{
		Item{Name: "Leather", Cost: 13, Dmg: 0, Def: 1},
		Item{Name: "Chainmail", Cost: 31, Dmg: 0, Def: 2},
		Item{Name: "Splitmail", Cost: 53, Dmg: 0, Def: 3},
		Item{Name: "Bandedmail", Cost: 75, Dmg: 0, Def: 4},
		Item{Name: "Platemail", Cost: 102, Dmg: 0, Def: 5},
	}

	ringShop := []Item{
		Item{Name: "Dmg +1", Cost: 25, Dmg: 1, Def: 0},
		Item{Name: "Dmg +2", Cost: 50, Dmg: 2, Def: 0},
		Item{Name: "Dmg +3", Cost: 100, Dmg: 3, Def: 0},
		Item{Name: "Def +1", Cost: 20, Dmg: 0, Def: 1},
		Item{Name: "Def +2", Cost: 40, Dmg: 0, Def: 2},
		Item{Name: "Def +3", Cost: 80, Dmg: 0, Def: 3},
	}

	min := math.MaxInt
	loadouts := generateCombinations(itemShop, armorShop, ringShop)

	for len(loadouts) > 0 {
		loadout := loadouts[0]
		loadouts = loadouts[1:]
		fmt.Println(loadout.TotalDmg, loadout.TotalDef)

		player := Player{HP: 100, Dmg: 0, Def: 0}
		boss := Player{HP: 109, Dmg: 8, Def: 2}
		for player.HP > 0 && boss.HP > 0 {
			fmt.Printf("\nPLAYER HP: %d, BOSS HP: %d\n", player.HP, boss.HP)
			totalAttack := player.Dmg + loadout.TotalDmg - boss.Def
			if totalAttack < 0 || boss.Def > totalAttack {
				totalAttack = 1
			}

			fmt.Printf("Player attacks! %d damage.\n", totalAttack)
			boss.HP -= totalAttack

			if boss.HP <= 0 {
				fmt.Println("You Won!")
				if loadout.TotalCost < min {
					min = loadout.TotalCost
				}
				break
			}

			totalAttack = boss.Dmg - (player.Def + loadout.TotalDef)
			if totalAttack < 0 || player.Def > totalAttack {
				totalAttack = 1
			}
			fmt.Printf("Boss attacks! %d damage.\n", totalAttack)
			player.HP -= totalAttack

			if player.HP <= 0 {
				fmt.Println("\nYou Lose!\n")
			}
		}
	}
	fmt.Println(min)
}
