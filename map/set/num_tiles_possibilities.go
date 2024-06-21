package data_structures

// clarifying questions
// * can it be the same letter? yes
// * seq can exclude some letters, but freq matters.
// * unique of the sequence matters
// * are tiles always sorted in order?

// approaches
// * actually generate all sequences - adv is can remove duplicates
// * some formula? probably not
// * if have a set of sequences for a set of tiles - how much the number of ways increases depends upon the uniqueness of the tile.. if it is completely new then have individual tile, then new tile at every possible location. if letter already exists then only half of insertions matter.

func numTilePossibilities(tiles string) int {
	setPoss := make(map[string]bool)

	for _, tile := range tiles {
		tileString := string(tile)
		newPoss := make(map[string]bool)
		newPoss[tileString] = true
		for sp, _ := range setPoss {
			insertTileIntoSeq(newPoss, sp, tileString)
		}
		for np, _ := range newPoss {
			setPoss[np] = true
		}
	}
	return len(setPoss)
}

func insertTileIntoSeq(newPoss map[string]bool, poss string, tile string) {
	for i := 0; i < len(poss); i++ {
		newSeq := string(poss[:i]) + tile + string(poss[i:])
		newPoss[newSeq] = true
	}
	newPoss[poss+tile] = true
}
