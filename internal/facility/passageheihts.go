package facility

//this file contains calculation for determining passage (ladder) heights

func (f *fac) calculatePassageHeights() {
	for levelIndex, _ := range f.undergroundFacility.levels {
		f.calculateLevelPassageHeights(levelIndex)
	}
}

func (f *fac) calculateLevelPassageHeights(levelIndex int) {
	var x float64
	for roomIndex, room := range f.undergroundFacility.levels[levelIndex].rooms {
		f.calculateRoomPassageHeights(levelIndex, roomIndex, x)
		x += room.width
	}
}

func (f *fac) calculateRoomPassageHeights(levelIndex, roomIndex int, x float64) {
	for passageIdx, _ := range f.undergroundFacility.levels[levelIndex].rooms[roomIndex].passages {
		f.calculateLowerPassageHeights(levelIndex, roomIndex, passageIdx, x)
	}

	for upperPassageIdx, _ := range f.undergroundFacility.levels[levelIndex].rooms[roomIndex].upperPassages {
		f.calculateUpperPassageHeights(levelIndex, roomIndex, upperPassageIdx, x)
	}
}

func (f *fac) calculateLowerPassageHeights(levelIndex, roomIndex, passageIdx int, x float64) {
	relatedLevelY := f.getNextLevelDimensionsAt(levelIndex, x)
	room := f.undergroundFacility.levels[levelIndex].rooms[roomIndex]
	f.undergroundFacility.levels[levelIndex].rooms[roomIndex].passages[passageIdx].ladderHeight = int(relatedLevelY - (f.undergroundFacility.levels[levelIndex].deepness + room.floor))
}

func (f *fac) calculateUpperPassageHeights(levelIndex, roomIndex, passageIdx int, x float64) {
	relatedLevelY := f.getPreviousLevelDimensionsAt(levelIndex, x)
	room := f.undergroundFacility.levels[levelIndex].rooms[roomIndex]
	f.undergroundFacility.levels[levelIndex].rooms[roomIndex].upperPassages[passageIdx].ladderHeight = int((f.undergroundFacility.levels[levelIndex].deepness + room.floor) - relatedLevelY)
}

func (f *fac) getNextLevelDimensionsAt(levelIndex int, x float64) float64 {
	if len(f.undergroundFacility.levels)-1 == levelIndex {
		return 0
	}

	nextLevelIndex := levelIndex + 1
	var x2 float64
	for _, room := range f.undergroundFacility.levels[nextLevelIndex].rooms {
		if x >= x2 && x <= x2+room.width {
			return f.undergroundFacility.levels[nextLevelIndex].deepness + room.ceil
		}
		x2 += room.width
	}

	return 0
}

func (f *fac) getPreviousLevelDimensionsAt(levelIndex int, x float64) float64 {
	if levelIndex == 0 {
		return 0
	}

	prevLevelIndex := levelIndex - 1

	var x2 float64
	for _, room := range f.undergroundFacility.levels[prevLevelIndex].rooms {
		if x >= x2 && x <= x2+room.width {
			return f.undergroundFacility.levels[prevLevelIndex].deepness + room.floor
		}
		x2 += room.width
	}

	return 0
}
