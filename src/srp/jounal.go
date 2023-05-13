package srp

import "strings"

type Journal struct {
	News []string
}

func (j *Journal) String() string {
	return strings.Join(j.News, "\n")
}

func (j *Journal) AddNews(news string) {
	// ....
}

func (j *Journal) RemoveNews(news string) {
	// ....
}

func (j *Journal) UpdateNews(news string) {
	// ....
}

func (j *Journal) DeleteNews(news string) {
	// ....
}
