package main;

import "io/ioutil";
import "strings";
import "strconv";
import "os/exec";
//import "fmt";

type CpuInfo struct {
    arch      string
	vendorId  string
	model     int32
	modelName string
	cores     int16
	mhz       float64
}

func getCpuInfo() CpuInfo {
    var kernelStat[] byte;
    var cpuStatLine []string
    var cores int16;
    var cpuInfo CpuInfo;

    kernelStat, err := ioutil.ReadFile("/proc/cpuinfo");

    cpuStatLine = strings.Split(string(kernelStat), "\n");

    if err != nil {
        return cpuInfo;
    }

    for _, line := range cpuStatLine {
        if (strings.HasPrefix(line, "processor\t: ") == true) {
            cores++;
        }

        if (strings.HasPrefix(line, "vendor_id\t: ") == true && cpuInfo.vendorId == "") {
            line = strings.Split(line, ":")[1];
            line = strings.TrimSpace(line);

            cpuInfo.vendorId = line;
        }
    
        if (strings.HasPrefix(line, "model\t\t: ") == true && cpuInfo.model == 0) {
            line = strings.Split(line, ":")[1];
            line = strings.TrimSpace(line);

            model, _ := strconv.ParseInt(line, 0, 32);

            cpuInfo.model = int32(model);
        }

        if (strings.HasPrefix(line, "model name\t: ") == true && cpuInfo.modelName == "") {
            line = strings.Split(line, ":")[1];
            line = strings.TrimSpace(line);

            cpuInfo.modelName = line;
        }
    }
    
    cpuInfo.cores = cores;

    kernelStat, err = exec.Command("/usr/bin/lscpu").Output();

    cpuStatLine = strings.Split(string(kernelStat), "\n");

    if err != nil {
        return cpuInfo;
    }

    for _, line := range cpuStatLine {
        if (strings.HasPrefix(line, "CPU max MHz:") == true && cpuInfo.mhz == 0) {
            line = strings.Split(line, ":")[1];
            line = strings.TrimSpace(line);
            line = strings.Replace(line, ",", ".", -1);

            mhz, _ := strconv.ParseFloat(line, 64);
            
            cpuInfo.mhz = mhz;
        }

        if (strings.HasPrefix(line, "Architecture:") == true && cpuInfo.arch == "") {
            line = strings.Split(line, ":")[1];
            line = strings.TrimSpace(line);
            
            cpuInfo.arch = line;
        }
    }

    return cpuInfo;
}

func cpuStats (cpuInfo CpuInfo) ([]int, string) {
    var kernelStat[] byte;
    var cpuStat []int;
    
    kernelStat, err := ioutil.ReadFile("/proc/stat");
    
    if err != nil {
        return cpuStat, err.Error();
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
            return cpuStat, err.Error();
        }
    }
    
    return cpuStat, "";
}
