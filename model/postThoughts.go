package model

//POST
func CreateThought(thought Thought) error {

	query = `insert into thoughts(title, content, author) values(?, ?, ?);`

	_, err := db.Exec(query, thought.Title, thought.Content, thought.Author)
	if err != nil {
		return err
	}

	return nil
}

func UpdateThought(thought Thought) error {
	query = `update thoughts set title=?, content=?, author=? where id=?;`

	_, err := db.Exec(query, thought.Title, thought.Content, thought.Author, thought.ID)
	if err != nil {
		return err
	}

	return nil
}
