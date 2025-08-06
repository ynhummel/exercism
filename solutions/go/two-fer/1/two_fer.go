package twofer

// ShareWith determines what will be said when you share the extra cookie.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
