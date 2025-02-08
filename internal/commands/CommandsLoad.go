package commands

func (c *Commands) Load() {
	toRegister := []CommandData{
		agg, login, register, reset, users,
	}
	for _, data := range toRegister {
		c.Register(data)
	}
}
