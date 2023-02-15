package model

func DeleteThoughts(id uint64) error {
	query = `delete from thoughts where id=?;`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
