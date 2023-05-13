package srp

import "io/ioutil"

func PrintNews(filename string, journal Journal) {
	_ = ioutil.WriteFile(filename, []byte(journal.String()), 0644)
}
