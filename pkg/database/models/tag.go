package models 

type Tag struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"unique;not null"`
    Tasks []Task `gorm:"many2many:task_tags;joinForeignKey:TagID;joinReferences:TaskID" json:"-"`
}
