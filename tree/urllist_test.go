package tree

func TestURLList_AddURL(t *testing.T) {
	l := NewURLList()
	l.AddURL("http://www.google.com")
}
