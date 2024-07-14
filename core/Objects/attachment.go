package Objects

type Attachment struct {
	ID       int32
	ItemID   int32
	FileType string
	FileName string
	FileSize int32
	FilePath string
}

func (attachment *Attachment) GetItem() {

}
func (attachment *Attachment) SetItem() {

}
func (attachment *Attachment) Delete() {}
