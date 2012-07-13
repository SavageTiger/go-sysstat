package main;

import "strconv";

func StringToInteger(s string) int {
    var integer int;

    integer, err := strconv.Atoi(s);
    
    if err != nil {
        return 0;
    }
    
    return integer;
}
