package main

import "fmt"

// var lock = &sync.Mutex{}
// var once sync.Once

// type single struct {
// }

// var singleInstance *single

// func getInstance() *single {

// 	if singleInstance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if singleInstance == nil {
// 			fmt.Println("Creating single instance now...")
// 			singleInstance = &single{}
// 		} else {
// 			fmt.Println("Single instance already created")
// 		}

// 	} else {
// 		fmt.Println("Single instance already created")
// 	}
// 	return singleInstance
// }

// func getInstance() *single {
// 	if singleInstance == nil {
// 		once.Do(
// 			func() {
// 				fmt.Println("Creating single instance now ...")
// 				singleInstance = &single{}
// 			})
// 	} else {
// 		fmt.Println("Single instance already created")
// 	}

// 	return singleInstance
// }

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}

type Ak47 struct {
	gun
}

func NewAK47() IGun {
	return &Ak47{
		gun: gun{
			name:  "AK47",
			power: 4,
		},
	}
}

type musket struct {
	gun
}

func NewMuskket() IGun {
	return &musket{
		gun: gun{
			name:  "Musket",
			power: 1,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "AK47" {
		return NewAK47(), nil
	} else if gunType == "Musket" {
		return NewMuskket(), nil
	}
	return nil, fmt.Errorf("wrong gun type")
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}

func main() {
	ak47, _ := getGun("AK47")
	musket, _ := getGun("Musket")

	printDetails(ak47)
	printDetails(musket)

}
