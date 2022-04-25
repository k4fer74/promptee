package lyrics

type Registry interface {
	GetSong(id int) (*Song, error)
	GetSongs() ([]*Song, error)
}
