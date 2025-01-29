package game

type player struct {
	id   int
	name string
}

func newPlayer(id int, name string) *player {
	return &player{
		id:   id,
		name: name,
	}
}
