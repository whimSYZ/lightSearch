package lightSearch

type pair struct {
    id      int
    val   float64
}

type token struct {
    str		string
    docCount    int
    indicies    map[int]pair
}