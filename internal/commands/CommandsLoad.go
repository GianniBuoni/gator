package commands

func (c *Commands) Load() {
	toRegister := []CommandData{
		addfeed, agg, feeds, login, register, reset, users,
	}
	for _, data := range toRegister {
		c.Register(data)
	}
}
