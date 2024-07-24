
### Context

Collect prometheus metrics for Linux machines to be consumed using Prometheus and Grafana.

### Prometheus & Grafana

* Following is the `compose.yaml` file for Docker compose

```
services:
  prometheus:
    image: prom/prometheus
    container_name: prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
    ports:
      - 9090:9090
    restart: unless-stopped
    volumes:
      - ./prometheus:/etc/prometheus
      - prom_data:/prometheus
    networks:
      - monitoring
  grafana:
    image: grafana/grafana
    container_name: grafana
    ports:
      - 3000:3000
    restart: unless-stopped
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=grafana
    volumes:
      - ./grafana:/etc/grafana/provisioning/datasources
    networks:
      - monitoring
volumes:
  prom_data:
networks:
  monitoring:
```

* Prometheus `prometheus.yml` config file

```
global:
  scrape_interval: 7s
  scrape_timeout: 5s
  evaluation_interval: 5s
alerting:
  alertmanagers:
  - static_configs:
    - targets: []
    scheme: http
    timeout: 10s
    api_version: v1
scrape_configs:
- job_name: host-node
  honor_timestamps: true
  scrape_interval: 5s
  scrape_timeout: 5s
  metrics_path: /metrics
  scheme: http
  static_configs:
  - targets:
    - 192.168.1.3:8880
```

* Grafana `datasource.yml` config file

```
apiVersion: 1

datasources:
- name: Prometheus
  type: prometheus
  url: http://prometheus:9090 
  isDefault: true
  access: proxy
  editable: true
```

### Visualization

Dashboard `.json` can be found [in this link ](https://grafana.com/grafana/dashboards/1860-node-exporter-full/)

### node-exporter

Dashboard utilize the `node-exporter` [open source project ](https://github.com/prometheus/node_exporter)

### Stress Test Results

| Disk Stress  | Comments                                                                          | Screenshots                                                                                                                                                                                                                                                                                                                                     |
|--------------|-----------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Disk Stress  | Command to stress test storage - **`stress-ng --hdd 5`**. 10mins running results  | ![disk.png](visualization%2F1%2Fdisk.png) ![memory.png](visualization%2F1%2Fmemory.png) ![process.png](visualization%2F1%2Fprocess.png) ![system.png](visualization%2F1%2Fsystem.png) ![pressure.png](visualization%2F1%2Fpressure.png) ![storage.png](visualization%2F1%2Fstorage.png) ![storage_ext.png](visualization%2F1%2Fstorage_ext.png) |
| CPU Stress  | Command to stress test CPU - **`stress-ng --cpu 4`**. 10mins running results  | ![summary.png](visualization%2F2%2Fsummary.png) ![basic_cpu.png](visualization%2F2%2Fbasic_cpu.png) ![cpu.png](visualization%2F2%2Fcpu.png)                                                                                                                                                                                                     |
| All Stress  | Command to stress all resources - **`stress-ng --cpu 4 --io 3 --vm 3 --vm-bytes 3G --timeout 600s`**. Test timeout 60s | ![disk_iops.png](visualization%2F3%2Fdisk_iops.png) ![disk_rw_stats.png](visualization%2F3%2Fdisk_rw_stats.png) ![io_utilization.png](visualization%2F3%2Fio_utilization.png) ![system_pressure.png](visualization%2F3%2Fsystem_pressure.png)                                                                                                   |
| Go memory stress  | The stress test is using code available in `src/4/main.go`. Compile and run the code and watch the result in the dashboard. | ![basic_cpu_mem_disk.png](visualization%2F4%2Fbasic_cpu_mem_disk.png) ![cpu_memory_stack.png](visualization%2F4%2Fcpu_memory_stack.png) ![memory_pages.png](visualization%2F4%2Fmemory_pages.png) ![stall_information.png](visualization%2F4%2Fstall_information.png) ![io_utilization.png](visualization%2F4%2Fio_utilization.png)             |

