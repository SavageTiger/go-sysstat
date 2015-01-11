package main

import "time";
import "strconv";
import "fmt"; // Debugging.

func monitorLoop() bool {
    var second time.Duration = (1000 * time.Millisecond);
    var cpuInfo CpuInfo = getCpuInfo();
    
    fmt.Println("CPU [" + cpuInfo.arch +"]: " + cpuInfo.modelName + " [" + strconv.Itoa(int(cpuInfo.cores)) + " cores, " + strconv.Itoa(int(cpuInfo.mhz)) + " mhz]");
    
    for ;; {
        fmt.Println(cpuStats(cpuInfo));
        
        time.Sleep(second * 1);
    }
    
    return false;
}

