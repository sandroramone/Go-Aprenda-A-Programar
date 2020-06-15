// - Partindo do código abaixo, utilize unmarshal e demonstre os valores.
//     - https://play.golang.org/p/b_UuCcZag9
// - Dica: JSON-to-Go.

package main

import (
	"encoding/json"
	"fmt"
)

type Persons []struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var x Persons
	json.Unmarshal([]byte(s), &x)

	for _, v := range x {
		fmt.Printf("\n%v %v %v\n", v.First, v.Last, v.Age)
		for _, v := range v.Sayings {
			fmt.Printf("\t%v\n", v)
		}
	}
}
