package pudl

func pudlsForTesting() (*[]Pudl, *[]Pudl) {
	oldPudls := &[]Pudl{
		{
			Title: "A",
			Url:   "https://lol.com/A",
		},
		{
			Title: "B",
			Url:   "https://lol.com/B",
		},
		{
			Title: "C",
			Url:   "https://lol.com/C",
		},
		{
			Title: "D",
			Url:   "https://lol.com/D",
		},
	}

	pudlsToShow := &[]Pudl{
		{
			Title: "A",
			Url:   "https://lol.com/A",
		},
		{
			Title: "J",
			Url:   "https://lol.com/J",
		},
		{
			Title: "O",
			Url:   "https://lol.com/O",
		},
	}
	return oldPudls, pudlsToShow
}

func ExampleDisplay() {

	oldPudls, pudlsToShow := pudlsForTesting()

	Display(pudlsToShow, oldPudls, -1)
	// Output: ⭐⭐⭐ 😹 Lista tematow z pudla 🐩 😹 ⭐⭐⭐
	//1: A
	//2:⭐ J
	//3:⭐ O

}
