package repository

//This repository must searching data of PG SQL

import (
	"database/sql"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/model"
)

type ImpReposityAlbumPg struct {
	db *sql.DB
}

//ADD ALBUM
func (pg *ImpReposityAlbumPg) AddAlbum(newAlbun model.Album) (model.Album, error) {

	var newAlbum model.Album
	query := pg.db.QueryRow("INSERT INTO album (title, artist, price) values ($1, $2, $3) returning *", newAlbun.Title, newAlbun.Artist, newAlbun.Price)

	if err := query.Scan(&newAlbum.ID, &newAlbum.Artist, &newAlbum.Title, &newAlbum.Price); err != nil {
		return newAlbum, err
	}

	return newAlbum, nil
}

//REMOVE ALBUM
func (pg *ImpReposityAlbumPg) RemoveAlbumById(idAlbum int64) (int64, error) {

	result, err := pg.db.Exec("DELETE FROM album WHERE id = $1", idAlbum)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, nil
}

//FIND ALL ALBUMS
func (pg *ImpReposityAlbumPg) Find() ([]model.Album, error) {

	rows, err := pg.db.Query("SELECT * FROM album")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var albums []model.Album
	for rows.Next() {
		var alb model.Album
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			return albums, err
		}
		albums = append(albums, alb)
	}

	if albums == nil {
		albums = []model.Album{}
	}

	if err := rows.Err(); err != nil {
		return albums, err
	}

	return albums, err

}

func ImpReposityPgContructor(dbSql *sql.DB) *ImpReposityAlbumPg {
	return &ImpReposityAlbumPg{db: dbSql}
}
