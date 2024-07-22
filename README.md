### Context

Collect prometheus metrics for Linux machines to be consumed using Prometheus and Grafana.

### Linux

This project https://github.com/flaviostutz/perfstat is used for collection Linux metrics.

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
    driver: bridge
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

The `.json` file for Grafana visualization is under `grafana` directory for using with `perfstat` project.

### node-exporter

Open source dashboard showing extensive metrics from `node-exporter` https://github.com/rfmoz/grafana-dashboards (via https://grafana.com/grafana/dashboards/1860-node-exporter-full/)