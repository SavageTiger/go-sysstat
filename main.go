/*

    This is a little google go project written just to experiment with go.
    
    Autor: Sven Hagemann, sven@savagetiger.org
    
*/

package main

import "fmt"

func main() {
    if monitorLoop() {
    	fmt.Println("It has been done");
    }
    
    fmt.Println("Ok");

    return;
}
