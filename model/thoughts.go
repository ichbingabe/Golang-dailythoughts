package model

type Thought struct {
	ID      uint64 `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllThoughts() ([]Thought, error) {
	var thoughts []Thought

	query := `select id, author, title, content from thoughts;`

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
