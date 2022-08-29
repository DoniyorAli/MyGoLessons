package main

func main(){ // STRING, NUMBER, BOOLEAN

// STRING

// o'zg - nomi - tipi  -  VALUE - qiymat 
	var message string = "Hello Student" 
	println(message)

	var message1 string        // DECLARE --> tanitdim 
	message1 = "Hello World"   // ASSIGNMEND --> tenglashtirdim
	println(message1)

	var box1, box2 string = "first", "second"
	println(box1,box2)

// NOTE !

	var box = "This is variable" // TIPI berilmagan ammo bu ham ishlaydi !
	println(box)                 // tipi berilmasa o'zi avtomat topib sekin ishlaydi 
			                // agar tipi berilsa kod tezlik jihatidan yaxshi ishlaydi  
//_______________________________________________________________________

// NUMBER
// INT
	var a,b,c int    // BUTUN SON  Ex: 7, -5, 11, 23, -94, 768, 811
	println(a,b,c)   // --> 0 0 0

	var n1, n2, n3 int = 11, 22, 33
	println(n1, n2, n3)

	var num1, num2, num3 int = 11, 22, 33
	println(num1 + num2 + num3)
//_______________________________________________________________________

// FLOAT

	var x, y, z float32  // QOLDIQ SON  Ex: 1.5, -11.6. 123.5, 4567.11, -12.5
	println(x,y,z)     // --> +0.000000e+000 +0.000000e+000 +0.000000e+000

	var f1, f2, f3 float64 = 11.1, 22.2, 33.3
	println(f1, f2, f3)

	var flo1, flo2, flo3 float64 = 11.1, 22.2, 33.3
	println(flo1 + flo2 + flo3)
//_______________________________________________________________________

// BOOLEAN

	var erkakMi bool = true  // false
	println(erkakMi)

//_______________________________________________________________________

// HAMMA O'ZGARUVCHILAR
// Yani bitta satrga togridan togri harhil tipda qiymat birlashtirish mumkun!

	var q, w, e, r = 11, 22.3, "Hello", true
	println(q, w, e, r)
//_______________________________________________________________________

// NEW NOTE !
//  := --> bu operator ishlatilsa VAR va TIP yozilishi shartamas !
//  o'zgaruvchisini va typini o'zi avtomat tanlab qo'yadi

	num, flo, str, truFal := 11, 22.3, "Hello", false
	println(num, flo, str, truFal)   // bir satrlik o'zgaruvchi typelar kodi !

	var bag1, bag2, bag3, bag4 = 33, 44.5, "Hello", true
	println(bag1, bag2, bag3, bag4)


	// vBox := 11
	// println(vBox)

	// vFlo := 22.3
	// println(vFlo)

	// vStr := "Hello!"
	// println(vStr)

	// vBool := true
	// println(vBool)
}