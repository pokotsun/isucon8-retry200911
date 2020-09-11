package main

func fetchSheetFromID(id int64) Sheet {
	switch {
	case (0 < id && id <= 50):
		return Sheet{ID: id, Rank: "S", Num: id, Price: 5000}
	case (50 < id && id <= 200):
		return Sheet{ID: id, Rank: "A", Num: id - 50, Price: 3000}
	case (200 < id && id <= 500):
		return Sheet{ID: id, Rank: "B", Num: id - 200, Price: 1000}
	case (500 < id && id <= 1000):
		return Sheet{ID: id, Rank: "C", Num: id - 500, Price: 0}
	default:
		return Sheet{ID: -1, Rank: "INVALID", Num: -1, Price: -1}
	}
}

func fetchAllSheets() []Sheet {
	sheets := make([]Sheet, 1000, 1000)
	for i := 0; i < 1000; i++ {
		sheets[i] = fetchSheetFromID(int64(i + 1))
	}
	return sheets
}

func fetchSheetIDList() []int64 {
	idList := []int64{}
	for i := 1; i <= 1000; i++ {
		idList = append(idList, int64(i))
	}
	return idList
}
