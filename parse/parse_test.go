package parse

import (
	"fmt"
)

func ExampleMain() {
	exp := "1+2/3*4-(5+6)/7-8+9"
	// exp := "1+2/3"
	p := Parse(exp)
	fmt.Println(*p)
	// Output:
	// {+:
	// 	1
	// 	{-:
	// 		{*:
	// 			{/:
	// 				2
	// 				3
	// 			}
	// 			4
	// 		}
	// 		{/:
	// 			{+:
	// 				5
	// 				6
	// 			}
	// 			7
	// 		}
	// 		8
	// 	}
	// 	9
	// }
}

func ExampleParseBySign() {
	exp := "1+2/3*4-(5+6)/7-8+9"
	ch := make(chan string)
	go parseBySign(exp, '+', ch )
	for p := range ch {
		fmt.Println(p)
	}
	// Output: 
	// 1
	// 2/3*4-(5+6)/7-8
	// 9
}
