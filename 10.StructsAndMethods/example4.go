package main

import "fmt"

//*========================STRUCT============================
type User struct {
	ID, Age                 int
	Name, Surname, Username string
	Pay                     *Payment
}

type Payment struct {
	Salary, Bonus float64   
}
//*========================CONSTRUCTOR=======================
func NewUser() *User {
	var n = new(User)      // n := User() or n := new(User)
	return n
}
//*========================METHODS===========================
func (u User) GetFullName() string {
	return u.Name + " " + u.Surname
}

func (p Payment) AddSalBon() float64 {
	return float64(p.Salary) + float64(p.Bonus)
}

func main() {

	u1 := User{ID: 1,
		Age: 27,
		Name: "Ali",
		Surname: "Khanov",
		Username: "AliKhan",
		Pay: &Payment{Salary: 2000, Bonus: 500},
	}

	fmt.Println(u1.Name)
	fmt.Println(u1.Pay.Salary)

	fmt.Println(NewUser())	// ! ---> nil

	fmt.Println(u1.GetFullName())

	fmt.Println(u1.Pay.AddSalBon())
}
