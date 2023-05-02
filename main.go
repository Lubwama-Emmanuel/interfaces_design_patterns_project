// package main

// import "fmt"

// // var lock = &sync.Mutex{}
// // var once sync.Once

// // type single struct {
// // }

// // var singleInstance *single

// // func getInstance() *single {

// // 	if singleInstance == nil {
// // 		lock.Lock()
// // 		defer lock.Unlock()
// // 		if singleInstance == nil {
// // 			fmt.Println("Creating single instance now...")
// // 			singleInstance = &single{}
// // 		} else {
// // 			fmt.Println("Single instance already created")
// // 		}

// // 	} else {
// // 		fmt.Println("Single instance already created")
// // 	}
// // 	return singleInstance
// // }

// // func getInstance() *single {
// // 	if singleInstance == nil {
// // 		once.Do(
// // 			func() {
// // 				fmt.Println("Creating single instance now ...")
// // 				singleInstance = &single{}
// // 			})
// // 	} else {
// // 		fmt.Println("Single instance already created")
// // 	}

// // 	return singleInstance
// // }

// type IGun interface {
// 	setName(name string)
// 	setPower(power int)
// 	getName() string
// 	getPower() int
// }

// type IDatabase interface {
// 	create(string, map[string]string) (bool, error)
// }

// type phone struct {
// 	info     map[string]string
// 	location string
// }

// func (p *phone) create(s string, m map[string]string) bool

// type gun struct {
// 	name  string
// 	power int
// }

// func (g *gun) setName(name string) {
// 	g.name = name
// }

// func (g *gun) setPower(power int) {
// 	g.power = power
// }

// func (g *gun) getName() string {
// 	return g.name
// }

// func (g *gun) getPower() int {
// 	return g.power
// }

// type Ak47 struct {
// 	gun
// }

// func NewAK47() IGun {
// 	return &Ak47{
// 		gun: gun{
// 			name:  "AK47",
// 			power: 4,
// 		},
// 	}
// }

// type musket struct {
// 	gun
// }

// func NewMuskket() IGun {
// 	return &musket{
// 		gun: gun{
// 			name:  "Musket",
// 			power: 1,
// 		},
// 	}
// }

// func getGun(gunType string) (IGun, error) {
// 	if gunType == "AK47" {
// 		return NewAK47(), nil
// 	} else if gunType == "Musket" {
// 		return NewMuskket(), nil
// 	}
// 	return nil, fmt.Errorf("wrong gun type")
// }

// func printDetails(g IGun) {
// 	fmt.Printf("Gun: %s", g.getName())
// 	fmt.Println()
// 	fmt.Printf("Power: %d", g.getPower())
// 	fmt.Println()
// }

// func main() {
// 	ak47, _ := getGun("AK47")
// 	musket, _ := getGun("Musket")

// 	printDetails(ak47)
// 	printDetails(musket)

// }
package main

import "fmt"

type Database interface {
	create(string, map[string]string)
	read(string) map[string]string
	update(string, map[string]string)
	delete(string)
}

type phone struct {
	location    string
	information map[string]string
}

func (p *phone) create(l string, s map[string]string) {
	p.location = l
	p.information = s
}

func (p phone) read(l string) map[string]string {
	if p.location == l {
		return p.information
	}
	return nil
}

func (p *phone) update(l string, s map[string]string) {
	if p.location == l {
		p.information = s
	}
}

func (p *phone) delete(l string) {
	if p.location == l {
		p = nil
	}
}

func newStruct() *phone {
	return &phone{}
}

func main() {
	new := newStruct()
	map1 := map[string]string{
		"Name":   "Emmanuel",
		"Second": "Lubwama",
	}
	new.create("first", map1)
	new.update("first", map[string]string{
		"Name":   "Rex",
		"Second": "Munil",
	})

	new.delete("first")
	value := new.read("first")

	fmt.Println(new)
	fmt.Println(value)

}
