package main

func main() {
	karn := person{
		firstName: "Nitipat",
		lastName:  "L",
		contact: contact{
			email: "iamsvz@gmail.com",
			zip:   50200,
		},
	}

	karn.print()
}
