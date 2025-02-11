package commands

func (c *Commands) Load() {
	toRegister := []CommandData{
		addfeed, agg,
		browse,
		feeds, follow, following,
		login,
		register, reset,
		unfollow, users,
	}
	for _, data := range toRegister {
		c.Register(data)
	}
}
