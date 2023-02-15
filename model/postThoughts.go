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
