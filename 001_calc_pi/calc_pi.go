package calcpi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math/big"
)

//logic/code used from https://play.golang.org/p/hF9jklt5lp
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many digits of Pi do you want? ")
	text, _ := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")

	digits, _ := strconv.ParseInt(text, 10, 64)

	pi := getPi(digits)

	fmt.Println(pi)
}

func getPi(numberofdigits int64) string {

	digits := big.NewInt(500+10)
	unity := big.NewInt(0)
	unity.Exp(big.NewInt(10), digits, nil)
	pi := big.NewInt(0)
	four := big.NewInt(4)
	pi.Mul(four, pi.Sub(pi.Mul(four, arccot(5, unity)), arccot(239, unity)))
	//pi := arccot(numberofdigits, unity)

	piString := pi.String()

	runes := []rune(piString)

	piReturn := string(runes[0:1]) + "." + string(runes[1:numberofdigits])

	return piReturn
}

func arccot(x int64, unity *big.Int) *big.Int {
	bigx := big.NewInt(x)
	xsquared := big.NewInt(x * x)
	sum := big.NewInt(0)
	sum.Div(unity, bigx)
	xpower := big.NewInt(0)
	xpower.Set(sum)
	n := int64(3)
	zero := big.NewInt(0)
	sign := false

	term := big.NewInt(0)
	for {
		xpower.Div(xpower, xsquared)
		term.Div(xpower, big.NewInt(n))
		if term.Cmp(zero) == 0 {
			break
		}
		if sign {
			sum.Add(sum, term)
		} else {
			sum.Sub(sum, term)
		}
		sign = !sign
		n += 2
	}
	return sum
}