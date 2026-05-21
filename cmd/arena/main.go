package main

import "fmt"

func calculateDamage(baseDamage int, bonusDamage int) int {
	return baseDamage + bonusDamage
}

func main() {
	fmt.Println("Welcome to DevArena")

	heroName := "Ragnar"
	heroClass := "Warrior"
	heroLevel := 1
	heroHP := 100
	heroBaseDamage := 15
	heroBonusDamage := 5
	heroAlive := true
	heroCriticalChance := 0.15

	enemyName := "Goblin"
	enemyHP := 60

	heroTotalDamage := calculateDamage(heroBaseDamage, heroBonusDamage)

	fmt.Println("Hero name:", heroName)
	fmt.Println("Hero class:", heroClass)
	fmt.Println("Hero level:", heroLevel)
	fmt.Println("Hero HP:", heroHP)
	fmt.Println("Hero base damage:", heroBaseDamage)
	fmt.Println("Hero bonus damage:", heroBonusDamage)
	fmt.Println("Hero total damage:", heroTotalDamage)
	fmt.Println("Hero alive:", heroAlive)
	fmt.Println("Hero critical chance:", heroCriticalChance)

	if heroHP > 0 {
		fmt.Println(heroName, "is ready to fight")
	} else {
		fmt.Println(heroName, "is defeated and cannot fight")
	}

	fmt.Println("Enemy name:", enemyName)
	fmt.Println("Enemy HP:", enemyHP)

	fmt.Println("Battle started")

	for round := 1; enemyHP > 0; round++ {
		enemyHP -= heroTotalDamage

		if enemyHP < 0 {
			enemyHP = 0
		}

		fmt.Println("Round:", round)
		fmt.Println(heroName, "hits", enemyName, "for", heroTotalDamage, "damage")
		fmt.Println(enemyName, "HP:", enemyHP)
	}

	fmt.Println("Battle finished")
	fmt.Println(enemyName, "is defeated")
}
