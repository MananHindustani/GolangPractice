package main

/*import (
	"fmt"
	"os"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

func main() {
	f, err := os.Open("/home/billion/pro.go")
	if err != nil {
		msg := err.Error()
		fmt.Println(msg)
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}*/
/*import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}*/

/*import (
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)
}*/
/*
import (
	"errors"
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}
func main() {
	radius := 20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f\n", area)
}
*/
import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f \n", area)
}
