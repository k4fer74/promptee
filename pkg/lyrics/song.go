package lyrics

type Song struct {
	ID     int      `json:"id"`
	Name   string   `yaml:"name" json:"name"`
	Author string   `yaml:"author" json:"author"`
	Lyrics []string `yaml:"lyrics" json:"lyrics"`
}
