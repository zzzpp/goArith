package main_test

import (
	"fmt"
	"github.com/zzzpp/goarith/parse"
)

func ExampleMain() {
	exp := "1+2/3*4-(5+6)/7-8+9"
	p := parse.Parse(exp)
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
