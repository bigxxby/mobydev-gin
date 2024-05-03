package genres

func (db *GenreRepository) UpdateGenre(id int, name string, description string) error {
	query := `
		UPDATE genres
		SET name = $2, description = $3
		WHERE id = $1
	`

	_, err := db.Database.Exec(query, id, name, description)
	if err != nil {
		return err
	}

	return nil
}
