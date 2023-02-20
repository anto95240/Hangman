package lstFunc

var (
	i int
)


func CreateProgressArray(word string) ([]string, []string) {
	res := make([]string, len(word)) // on ajoute un tableau de string et la longueur de word à res
	letters := []string{}
	for _, v := range word { // on parcours le mot
		letters = append(letters, string(v)) //on ajoute letters et le string de v a la variable letters
	}

	for i := range res { // on parcours la longueur de res
		if word[i] != ' ' { // si res d'indice i est différent d'une rune vide
			res[i] = "_" // on remplace cette par "_"
		} else { // skip space
			res[i] = " " // sinon on remplace cette lettre par un espace
		}
	}

	return res, letters

}
