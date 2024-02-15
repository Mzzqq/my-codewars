package main

func HowMuchILoveYou(i int) string {
	if i == 6 || i%6 == 0 {
		return "not at all"
	} else {
		switch res := i % 6; res {
		case 1:
			return "I love you"
		case 2:
			return "a little"
		case 3:
			return "a lot"
		case 4:
			return "passionately"
		case 5:
			return "madly"
		}
	}
	return ""
}
