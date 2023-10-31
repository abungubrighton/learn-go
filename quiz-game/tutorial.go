package main
import "fmt"

func main(){
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scan(&name)
	fmt.Printf("Hello %v and welcome to the quiz game\n",name)
	fmt.Printf("Your name is: %s\n", name)

	fmt.Println("Please enter age:")
	var age uint
	fmt.Scan(&age)

	if age>=10 {
		fmt.Println("Yeah you can play")
	}else{
		fmt.Println("You cannot Play")
		return
	}

	fmt.Printf("What is better ,the RTX 3080 or RTX 3090?:")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	fmt.Println(answer, answer2)
	if answer + " " + answer2  == "RTX 3090"{
		fmt.Println("Correct answer")

	}else if answer + " " + answer2 == "rtx 3090"{
		println("Correct answer")
	}else{
		fmt.Println("Incorrect")
	}
}