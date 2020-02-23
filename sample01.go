package sample

func HelloWorld(s string) string {
	base := "Hello World"
	if s == "" {
		return base + "!"
	} else {
		return base + ", " + s + "!"
	}
}
