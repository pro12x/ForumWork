package internal

import (
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

func Insertion(username, email, mdp string) bool {
	mdpHash := Encryptage(mdp)
	database, err := sql.Open("sqlite3", "Forum.sqlite")
	if err != nil {
		log.Fatal("Error: Impossible de se connecter à la base de données")
	}
	defer database.Close()
	if Existe(database, username, email) {
		fmt.Println("This Information existed !!!")
		return true
	}

	statement, err := database.Prepare("INSERT INTO User(name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal("Error: Impossible de préparer la requête d'insertion")
	}
	defer statement.Close()

	_, err = statement.Exec(username, email, mdpHash)
	if err != nil {
		log.Fatal("Error: Impossible d'insérer les données dans la base de données")
	}

	fmt.Println("Utilisateur inséré avec succès.")
	return false
}

func Encryptage(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	Hasher := hash.Sum(nil)
	return hex.EncodeToString(Hasher)
}

func Existe(database *sql.DB, username, useremail string) bool {
	// Lecture des utilisateurs depuis la base de données
	rows, err := database.Query("SELECT id, name, email FROM User")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatal("Error: Impossible de scanner les résultats")
		}
		if strings.EqualFold(name, username) || strings.EqualFold(email, useremail) {
			return true
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return false
}

func Connexion(username, useremail, mdp string) bool {
	mdphash := Encryptage(mdp)
	database, err := sql.Open("sqlite3", "Forum.sqlite")
	if err != nil {
		log.Fatal("Error: Impossible de se connecter à la base de données")
	}
	defer database.Close()
	// Lecture des utilisateurs depuis la base de données
	rows, err := database.Query("SELECT name, email, password FROM User")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, email, password string
		if err := rows.Scan(&name, &email, &password); err != nil {
			log.Fatal("Error: Impossible de scanner les résultats")
		}
		if strings.EqualFold(name, username) || strings.EqualFold(email, useremail) {
			if subtle.ConstantTimeCompare([]byte(mdphash), []byte(password)) == 1 {
				return true
			}
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return false
}
