package model

//GET
func GetAllThoughts() ([]Thought, error) {
	var thoughts []Thought

	query = `select id, author, title, content from thoughts;`

	rows, err := db.Query(query)
	if err != nil {
		return thoughts, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var author, title, content string

		err := rows.Scan(&id, &author, &title, &content)
		if err != nil {
			return thoughts, err
		}

		thought := Thought{
			ID:      id,
			Author:  author,
			Title:   title,
			Content: content,
		}

		thoughts = append(thoughts, thought)
	}

	return thoughts, nil
}

func GetThoughtById(id uint64) (Thought, error) {
	var thought Thought

	query = `select id, title, content, author from thoughts where id=?`

	rows, err := db.Query(query, id)
	if err != nil {
		return thought, err
	}

	defer rows.Close()

	if rows.Next() {
		var id uint64
		var title, content, author string

		err := rows.Scan(&id, &title, &content, &author)
		if err != nil {
			return thought, err
		}

		thought = Thought{
			ID:      id,
			Title:   title,
			Content: content,
			Author:  author,
		}
	}

	return thought, nil
}

func GetThoughtByTitle(title string) ([]Thought, error) {
	var thoughts []Thought

	query = `select id, title, content, author from thoughts where title=(?)`

	rows, err := db.Query(query, title)
	if err != nil {
		return thoughts, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var title, content, author string

		err := rows.Scan(&id, &title, &content, &author)
		if err != nil {
			return thoughts, err
		}

		thought := Thought{
			ID:      id,
			Title:   title,
			Content: content,
			Author:  author,
		}

		thoughts = append(thoughts, thought)
	}
	return thoughts, nil
}
