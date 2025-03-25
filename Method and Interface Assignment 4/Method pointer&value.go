// Go program to illustrate how the method can accept pointer and value


package main
import "fmt"
type author struct {
name string
branch string
}
func (a *author) show_1(abranch string) {
(*a).branch = abranch
}

func (a author) show_2() {
	a.name = "ABC"
	fmt.Println("Author's name(Before) : ", a.name)
	}
	// Main function
	func main() {

	res := author{
	name: "XYZ",
	branch: "CSE",
	}
	fmt.Println("Branch Name(Before): ", res.branch)
	res.show_1("ECE")
	fmt.Println("Branch Name(After): ", res.branch)
(&res).show_2()
fmt.Println("Author's name(After): ", res.name)
}