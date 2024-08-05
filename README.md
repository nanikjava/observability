
### Repository

This repository contain experiment result to learn and understand pattern how Linux system react when it is under load in different scenarios and use cases. 

System specification used for this experiment:

```
HP EliteDesk 800 G2 SFF 
Intel(R) Core(TM) i5-6500 CPU @ 3.20GHz
8GiB System Memory
128GB SanDisk 
```

![elitedesk.jpeg](images%2Felitedesk.jpeg)

This repository use Prometheus as metric collector and Grafana to display the metric information. Out-of-the-box storage is used, so no additional storage are
used for storing the metrics.

### Goals

The end goal of this experiment is to learn as much as possible to understand different scenarios and use cases when system are under stress. This will give better
understanding regardless whether it is a single computer or mega cluster.

Understanding the different aspects of the hardware being stress tested will surface patterns that can be used to troubleshoot production environment. 

### Docker Compose

Docker compose file is available under [compose.yaml](compose.yaml)

### Prometheus & Grafana
* Prometheus [configuration file](prometheus%2Fprometheus.yml)

* Grafana [configuration file](grafana%2Fdatasource.yml)

### Visualization

Dashboard `.json` can be found [in this link ](grafana%2Fdashboard.json). The dashboard utilize the `node-exporter`
[open source project ](https://github.com/prometheus/node_exporter)

### Stress Test Results

| Disk Stress      | Comments                                                                                                                                                                                                                                                                        | Screenshots                                                                                                                                                                                                                                                                                                                                     |
|------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Disk Stress      | Command to stress test storage - **`stress-ng --hdd 5`**. 10mins running results                                                                                                                                                                                                | ![disk.png](visualization%2F1%2Fdisk.png) ![memory.png](visualization%2F1%2Fmemory.png) ![process.png](visualization%2F1%2Fprocess.png) ![system.png](visualization%2F1%2Fsystem.png) ![pressure.png](visualization%2F1%2Fpressure.png) ![storage.png](visualization%2F1%2Fstorage.png) ![storage_ext.png](visualization%2F1%2Fstorage_ext.png) |
| CPU Stress       | Command to stress test CPU - **`stress-ng --cpu 4`**. 10mins running results                                                                                                                                                                                                    | ![summary.png](visualization%2F2%2Fsummary.png) ![basic_cpu.png](visualization%2F2%2Fbasic_cpu.png) ![cpu.png](visualization%2F2%2Fcpu.png)                                                                                                                                                                                                     |
| All Stress       | Command to stress all resources - **`stress-ng --cpu 4 --io 3 --vm 3 --vm-bytes 3G --timeout 600s`**. Test timeout 60s                                                                                                                                                          | ![disk_iops.png](visualization%2F3%2Fdisk_iops.png) ![disk_rw_stats.png](visualization%2F3%2Fdisk_rw_stats.png) ![io_utilization.png](visualization%2F3%2Fio_utilization.png) ![system_pressure.png](visualization%2F3%2Fsystem_pressure.png)                                                                                                   |
| Go memory stress | The stress test is using code available in `src/4/main.go`. Compile and run the code and watch the result in the dashboard.                                                                                                                                                     | ![basic_cpu_mem_disk.png](visualization%2F4%2Fbasic_cpu_mem_disk.png) ![cpu_memory_stack.png](visualization%2F4%2Fcpu_memory_stack.png) ![memory_pages.png](visualization%2F4%2Fmemory_pages.png) ![stall_information.png](visualization%2F4%2Fstall_information.png) ![io_utilization.png](visualization%2F4%2Fio_utilization.png)             |
| API Errors       | This stress test uses `stress-ng` with job file [network.job](stress-ng-jobs/network.job) is to test how an API application in `src/api` reacts when the network `epoll` and `sock` stressor are enabled. Error<br/>generated can be seen in [errors.md](src/api/errors.md) | ![network.png](visualization%2Fapi%2Fnetwork.png) ![api_errors.png](visualization%2Fapi%2Fapi_errors.png) ![cpu.png](visualization%2Fapi%2Fcpu.png)                                                                                                                                                                                             |



