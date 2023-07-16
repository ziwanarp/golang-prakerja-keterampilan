package models

type Film struct { 
	Id int `json:"id"` 
	Judul string `json:"judul"`
	Genre string `json:"genre"`
	Tahun string `json:"tahun"`
}
