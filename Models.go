package main
// Like table: Unique ID, contains who likes who
type like struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	Who_likes    int  `json:"who_likes"`
	Who_is_liked int  `json:"who_is_liked"`
}
// User entity
type User struct {
	Id       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Location float32 `json:"location"`
	Gender   string  `json:"gender"`
	Email    string  `json:"email"`
}
// Get POST Json for API type getsuers
type getuser struct {
	Id       uint    `json:"id"`
	Distance float32 `json:"distance"`
}
// Helper struct to find matches
type helper struct {
	Id1 uint
	Id2 uint
}
