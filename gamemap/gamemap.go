package gamemap

type Map struct {
	Grid [100][4]byte
}

func NewMap() Map {
	var g [100][4]byte
	for i := 1; i < len(g); i++ {
		for j := 1; j < len(g[0]); j++ {
			g[i][j] = ' '
		}
	}
	return Map{
		Grid: g}
}
