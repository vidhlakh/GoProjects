package main
import(
    "fmt"
    "math"
    "net/http"
)
func doubleValue(values [] int, i int){
    values[i]*=4
}
func double(val int){
    val*=10
}
//PArameter pointer
func doublePtr(n *int){
    *n*=5
}
//return 
func sqroot(n float64) (float64,error){
    if n<0{
        return 0.0, fmt.Errorf("Error of Sqrt of negative val %f",n)
    }
    return math.Sqrt(n),nil
}

// Garbage collection
// make use of defer .. Defer call in reverse order
func cleanup(name string){
    fmt.Printf("Cleaning up %s\n",name)
}
func worker(){
    defer cleanup("A")
    defer cleanup("B")
}

// Challenge to Get URL and return Content Type Response HTTP Header
func contentType(url string)(string,error){
    resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() // Make sure we close the body

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" { // Return error if Content-Type header not found
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return ctype, nil
}   