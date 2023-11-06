package main

import "fmt"

// Factory pattern - making weapons
// Interface for weapon
type IWeapon interface {
	getType() string
	getName() string
	getAtk() int
	getMp() int
}

// Concrete weapon
type Weapon struct {
	weaponType string
	name       string
	atk        int
	mp         int
}

func (w Weapon) getType() string {
	return w.weaponType
}

func (w Weapon) getName() string {
	return w.name
}

func (w Weapon) getAtk() int {
	return w.atk
}

func (w Weapon) getMp() int {
	return w.mp
}

// Concrete products
type Mace struct {
	Weapon
}

func makeMace() IWeapon {
	var n string
	fmt.Print("Choose name for your weapon: ")
	fmt.Scanln(&n)
	return &Mace{
		Weapon{
			weaponType: "Mace",
			name:       n,
			atk:        20,
			mp:         0,
		},
	}
}

type Sword struct {
	Weapon
}

func makeSword() IWeapon {
	var n string
	fmt.Print("Choose name for your weapon: ")
	fmt.Scanln(&n)
	return &Sword{
		Weapon{
			weaponType: "Sword",
			name:       n,
			atk:        10,
			mp:         5,
		},
	}
}

type Dagger struct {
	Weapon
}

func makeDagger() IWeapon {
	var n string
	fmt.Print("Choose name for your weapon: ")
	fmt.Scanln(&n)
	return &Dagger{
		Weapon{
			weaponType: "Dagger",
			name:       n,
			atk:        5,
			mp:         15,
		},
	}
}

// Decorator pattern - upgrading weapons
// Decorator for atkack
type Sharper struct {
	w IWeapon
}

func (s Sharper) getType() string {
	return s.w.getType()
}

func (s Sharper) getName() string {
	return s.w.getName()
}

func (s Sharper) getAtk() int {
	return s.w.getAtk() + 5
}

func (s Sharper) getMp() int {
	return s.w.getMp()
}

// Decorator for Magic power
type Enchanter struct {
	w IWeapon
}

func (e Enchanter) getType() string {
	return e.w.getType()
}

func (e Enchanter) getName() string {
	return e.w.getName()
}

func (e Enchanter) getAtk() int {
	return e.w.getAtk()
}

func (e Enchanter) getMp() int {
	return e.w.getMp() + 5
}

// Facade pattern
// Facade object
type forgeFacade struct {
	inv []IWeapon
}

// Singleton of Facade
var singleInstance *forgeFacade

func getInstance() *forgeFacade {
	if singleInstance == nil {
		singleInstance = &forgeFacade{}
	}
	return singleInstance
}

// 1. Accessing inventory
func (f *forgeFacade) inventory() bool {
	if len(f.inv) != 0 {
		fmt.Println("╔═════════════════════╗")
		for i, w := range f.inv {
			fmt.Printf("------ %d ------ \nType: %s \nName: %s \nAttack: %d \nMP: %d\n", i+1, w.getType(), w.getName(), w.getAtk(), w.getMp())
		}
		fmt.Println("╚═════════════════════╝")
		return true
	} else {
		fmt.Println("╔═════════════════════╗")
		fmt.Println("Nothing here... \nConsider ordering something?")
		fmt.Println("╚═════════════════════╝")
		return false
	}
}

// 2. Making a weapon
func (f *forgeFacade) makeWeapon() {
	fmt.Println("╔═════════════════════╗")
	fmt.Println("Sir, what weapon you want? \n1. Mace \n2. Sword \n3. Dagger")
	fmt.Println("╚═════════════════════╝")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		f.inv = append(f.inv, makeMace())
		fmt.Println("╔═════════════════════╗")
		fmt.Println("Done! \nCheck your inventory")
		fmt.Println("╚═════════════════════╝")
	case 2:
		f.inv = append(f.inv, makeSword())
		fmt.Println("╔═════════════════════╗")
		fmt.Println("Done! \nCheck your inventory")
		fmt.Println("╚═════════════════════╝")
	case 3:
		f.inv = append(f.inv, makeDagger())
		fmt.Println("╔═════════════════════╗")
		fmt.Println("Done! \nCheck your inventory")
		fmt.Println("╚═════════════════════╝")
	default:
		fmt.Println("╔═════════════════════╗")
		fmt.Println("Please, \nchoose from our assortment")
		fmt.Println("╚═════════════════════╝")
		f.makeWeapon()
	}
}

// 3. Upgrading a weapon
func (f *forgeFacade) upgradeWeapon() {
	notEmptyInv := f.inventory()
	if notEmptyInv {
		fmt.Println("╔═════════════════════╗")
		fmt.Println("What weapon needs upgrade?")
		fmt.Println("╚═════════════════════╝")
		var i int
		fmt.Scan(&i)
		fmt.Println("╔═════════════════════╗")
		fmt.Println("What kind of upgrade? \n1. Sharpen \n2. Enchant")
		fmt.Println("╚═════════════════════╝")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			f.inv[i-1] = Sharper{f.inv[i-1]}
			fmt.Println("╔═════════════════════╗")
			fmt.Println("Done! \nNow your weapon is sharper")
			fmt.Println("╚═════════════════════╝")
		case 2:
			f.inv[i-1] = Enchanter{f.inv[i-1]}
			fmt.Println("╔═════════════════════╗")
			fmt.Println("Done! \nNow your weapon is enchanted")
			fmt.Println("╚═════════════════════╝")
		default:
			fmt.Println("╔═════════════════════╗")
			fmt.Println("Please, choose from our assortment")
			fmt.Println("╚═════════════════════╝")
			f.upgradeWeapon()
		}
	}
}

// Facade itself
func (f *forgeFacade) blacksmithFacade() {
LOOP:
	for {
		fmt.Println("╔═════════════════════╗")
		fmt.Println("Welcome to our blacksmith! \nWhat can i help you with? \n1. My inventory \n2. Make me a weapon \n3. Upgrade me a weapon \n4. Leave")
		fmt.Println("╚═════════════════════╝")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			f.inventory()
		case 2:
			f.makeWeapon()
		case 3:
			f.upgradeWeapon()
		case 4:
			fmt.Println("╔═════════════════════╗")
			fmt.Println("Good luck! Come back later")
			fmt.Println("╚═════════════════════╝")
			break LOOP
		default:
			fmt.Println("╔═════════════════════╗")
			fmt.Println("We don't offer such services")
			fmt.Println("╚═════════════════════╝")
		}
	}
}

func main() {
	init := getInstance()
	init.blacksmithFacade()
}
