How to setup prometheus


-> docker volume create prometheus-data

-> docker run -p 9090:9090 -v ./external/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml -v prometheus-data:/prometheus prom/prometheus