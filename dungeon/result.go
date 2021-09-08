package dungeon

type DungeonClear struct {
	Name  string
	Count int
}

type DungeonList struct {
	Year    int
	Month   int
	Total   int
	Ranking []DungeonClear
}
