package main

type T1000 struct {
	On    bool
	Ammo  int
	Power int
}

func (c *T1000) Shoot() bool {
	if c.On && c.Ammo > 0 {
		c.Ammo--
		return true
	}
	return false
}

func (c *T1000) RideBike() bool {
	if c.On && c.Power > 0 {
		c.Power--
		return true
	}
	return false
}

func main() {
	robot := T1000{Ammo: 30, Power: 100}
	testStruct := &robot
}
