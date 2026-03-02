package models

type ConnectionString struct {
	User     string
	DBName   string
	Password string
	SslMode  string
	Port     string
}

type Author struct {
	AuthorName string `json:"author_name"`
	AuthorID   int    `json:"author_id"`
}

type Books struct {
	BookName string `json:"book_name"`
	BookID   int    `json:"book_id"`
	AuthorID int    `json:"author_id"`
	Author   Author `json:"author" gorm:"foreignKey:AuthorID;references:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
