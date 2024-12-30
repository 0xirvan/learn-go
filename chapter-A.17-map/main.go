package main

func main() {
	// Map berisi pasangan key-value
	var scores = map[string]int{
		"John": 12,
		"Jane": 23,
		"Jack": 34,
	}

	// Iterasi data map
	for k, v := range scores {
		println(k, v)
	}

	// Menghapus data dari map
	delete(scores, "John")
	println("After deleting John:")
	for k, v := range scores {
		println(k, v)
	}

	// Mengecek apakah key ada di map
	if _, isExist := scores["Jane"]; isExist {
		println("Jane exists")
	} else {
		println("Jane doesn't exist")
	}

	// Menambahkan data ke map
	scores["Doe"] = 56
	println("After adding Doe:")
	for k, v := range scores {
		println(k, v)
	}

	// Kombinasi map dan slice
	var students = []map[string]string{
		{"name": "John", "score": "A"},
		{"name": "Jane", "score": "B"},
		{"name": "Jack", "score": "C"},
	}

	for _, student := range students {
		println(student["name"], student["score"])
	}

	// Membuat map kosong
	var emptyMap = map[string]string{}
	println(emptyMap)

}
