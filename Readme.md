# goarith
An exploratory alternative paser for arithmetic expression with golang.

## usage
```
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
```

## target
- assuming their are services only knows one particular arithmetic operation, this program tries to generate tasks that these services can solve.
- support multicasting

## todo
- verify expression before processing
- unify operations
