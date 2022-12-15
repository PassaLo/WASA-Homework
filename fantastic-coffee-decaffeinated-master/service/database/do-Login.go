package database

// Login the user by their username
func (db *appdbimpl) LoginUser(username string) (int64, error) {
	//Id of the user that is logging in
	var id int64
	//Query for the login
	//First a query to check if username already exists. If it exists, return the id.
	row, err := db.c.Query("SELECT id FROM users WHERE ? == username", username)
	if err != nil {
		return id, err
	}

	//Check if the query have found the user
	if !row.Next() {
		//If not make another query to add the user. We first have to prepare the user to add.
		var u User
		u.username = username

		res, err := db.c.Exec(`INSERT INTO users (username, follows, followers, banned, photos) VALUES (?, ?, ?, ?, ?)`,
			u.username, u.follows, u.followers, u.banned, u.photos)
		if err != nil {
			return id, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return id, err
		}
	} else {
		//If the user already exists scan the result of the query to get the id
		row.Scan(id)
	}
	return id, nil
}