package Models

type Whislist struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func (b *Whislist) TableName() string {
	return "whistlist"
}
