package main

import "time";
import "io/ioutil";
import "strings";
import "fmt"; // Debugging.

func monitorLoop() bool {
    var second time.Duration = (1000 * time.Millisecond);
    
    for ;; {
        fmt.Println(cpuStats());
        
        time.Sleep(second * 1);
    }
    
    return false;
}

func cpuStats () string {
    var kernelStat[] byte;
    var cpuStat string;
    
    kernelStat, err := ioutil.ReadFile("/proc/stat");
    
    if err != nil {
        return err.Error();
    } else {
        var cpuStatLine []string
        var percentage int;
        var idleFactor int;
        
        cpuStatLine = strings.Split(string(kernelStat), "\n");
        cpuStatLine = strings.Split(cpuStatLine[0], " ");
        
        percentage =
            StringToInteger(cpuStatLine[2]) +
            StringToInteger(cpuStatLine[3]) +
            StringToInteger(cpuStatLine[4]);
        idleFactor = StringToInteger(cpuStatLine[5]);
        
        percentage = percentage / (percentage + idleFactor);
        
        if err != nil {
            return err.Error();
        }
        
        fmt.Println(idleFactor);
    }
    
    return cpuStat;
}
