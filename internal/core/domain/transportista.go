package domain

type Transportista struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"id"`
	Nombre   string `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Telefono string `gorm:"type:varchar(15);not null;column:telefono" json:"telefono"`
	Email    string `gorm:"type:varchar(100);not null;unique;column:email" json:"email"`
	Activo   bool   `gorm:"default:true;column:activo" json:"activo"`
}
