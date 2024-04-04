package main

type Pokemon struct {
	Name  string
	Power string
}

func (p *Pokemon) Rename(name string) *Pokemon {
	p.Name = name
	return p
}

func (p *Pokemon) LearnPower(power string) *Pokemon {
	p.Power = power
	return p
}

func main() {
	p1 := Pokemon{"Charmander", "Flame"}
	p1.Rename("Charmeleon")
	p1.LearnPower("Super Flame")

	p2 := Pokemon{"Pikachu", "Bolt"}
	p2.Rename("Raichu").LearnPower("Super Bolt")
}
