package main

import "fmt"

const (
	defaultHeroClass          = "Warrior"
	defaultHeroLevel          = 1
	defaultHeroHP             = 100
	defaultHeroBaseDamage     = 15
	defaultHeroBonusDamage    = 5
	defaultHeroCriticalChance = 0.15

	defaultEnemyName = "Goblin"
	defaultEnemyHP   = 60
)

type Hero struct {
	Name           string
	Class          string
	Level          int
	HP             int
	BaseDamage     int
	BonusDamage    int
	Alive          bool
	CriticalChance float64
	Attacks        [3]string
	Inventory      []string
	Stats          map[string]int
}

type Enemy struct {
	Name string
	HP   int
}

type Battle struct {
	Hero  Hero
	Enemy Enemy
	Round int
}

func calculateDamage(baseDamage int, bonusDamage int) int {
	return baseDamage + bonusDamage
}

func main() {
	fmt.Println("Welcome to DevArena")

	hero := Hero{
		Name:           "Ragnar",
		Class:          defaultHeroClass,
		Level:          defaultHeroLevel,
		HP:             defaultHeroHP,
		BaseDamage:     defaultHeroBaseDamage,
		BonusDamage:    defaultHeroBonusDamage,
		Alive:          true,
		CriticalChance: defaultHeroCriticalChance,
		Attacks:        [3]string{"Slash", "Pierce", "Heavy Strike"},
		Inventory:      []string{"Small Potion", "Wooden Shield"},
		Stats: map[string]int{
			"strength": 10,
			"agility":  7,
			"stamina":  12,
		},
	}

	enemy := Enemy{
		Name: defaultEnemyName,
		HP:   defaultEnemyHP,
	}

	battle := Battle{
		Hero:  hero,
		Enemy: enemy,
		Round: 0,
	}

	heroTotalDamage := calculateDamage(hero.BaseDamage, hero.BonusDamage)

	inventoryBeforeBattle := make([]string, len(hero.Inventory))
	copy(inventoryBeforeBattle, hero.Inventory)

	fmt.Println("Hero name:", hero.Name)
	fmt.Println("Hero class:", hero.Class)
	fmt.Println("Hero level:", hero.Level)
	fmt.Println("Hero HP:", hero.HP)
	fmt.Println("Hero base damage:", hero.BaseDamage)
	fmt.Println("Hero bonus damage:", hero.BonusDamage)
	fmt.Println("Hero total damage:", heroTotalDamage)
	fmt.Println("Hero alive:", hero.Alive)
	fmt.Println("Hero critical chance:", hero.CriticalChance)

	fmt.Println("Hero stats:")
	for statName, statValue := range hero.Stats {
		fmt.Println(statName+":", statValue)
	}

	intellect, exists := hero.Stats["intellect"]
	if exists {
		fmt.Println("Intellect:", intellect)
	} else {
		fmt.Println("Intellect stat is not defined")
	}

	hero.Stats["strength"] = hero.Stats["strength"] + 2
	hero.Stats["intellect"] = 3

	fmt.Println("Hero stats after training:")
	for statName, statValue := range hero.Stats {
		fmt.Println(statName+":", statValue)
	}

	delete(hero.Stats, "intellect")

	fmt.Println("Hero stats after removing temporary intellect bonus:")
	for statName, statValue := range hero.Stats {
		fmt.Println(statName+":", statValue)
	}

	fmt.Println("Hero attacks:")
	for index, attack := range hero.Attacks {
		fmt.Println("Attack", index+1, ":", attack)
	}

	fmt.Println("Hero inventory before battle:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length before reward:", len(hero.Inventory))
	fmt.Println("Inventory capacity before reward:", cap(hero.Inventory))

	hero.Inventory = append(hero.Inventory, "Rusty Sword")

	fmt.Println("Hero inventory after reward:")
	for index, item := range hero.Inventory {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory copy before battle still:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length after reward:", len(hero.Inventory))
	fmt.Println("Inventory capacity after reward:", cap(hero.Inventory))

	if hero.HP > 0 {
		fmt.Println(hero.Name, "is ready to fight")
	} else {
		fmt.Println(hero.Name, "is defeated and cannot fight")
	}

	fmt.Println("Enemy name:", enemy.Name)
	fmt.Println("Enemy HP:", enemy.HP)

	fmt.Println("Battle started")

	for battle.Enemy.HP > 0 {
		battle.Round++

		attackIndex := (battle.Round - 1) % len(battle.Hero.Attacks)
		selectedAttack := battle.Hero.Attacks[attackIndex]

		battle.Enemy.HP -= heroTotalDamage

		if battle.Enemy.HP < 0 {
			battle.Enemy.HP = 0
		}

		fmt.Println("Round:", battle.Round)
		fmt.Println(battle.Hero.Name, "uses", selectedAttack)
		fmt.Println(battle.Hero.Name, "hits", battle.Enemy.Name, "for", heroTotalDamage, "damage")
		fmt.Println(battle.Enemy.Name, "HP:", battle.Enemy.HP)
	}

	fmt.Println("Battle finished")
	fmt.Println(battle.Enemy.Name, "is defeated")
	fmt.Println(battle.Hero.Name, "received item:", "Rusty Sword")
}
