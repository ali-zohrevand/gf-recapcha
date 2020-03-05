package models

type Conf struct {
	Title string  `yaml:"title"`
	V3    Key     `yaml:"v3"`
	V2    Key     `yaml:"v2"`
	Score float32 `yaml:"score"`
}

type Key struct {
	Public  string `yaml:"public"`
	Private string `yaml:"private"`
}
