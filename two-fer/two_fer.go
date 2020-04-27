// Two-fer solution
package twofer

/* Returns one for you or given name */
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for "+ name + ", one for me."
}
