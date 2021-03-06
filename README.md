<div align="center">
  <img src="./temperature-sensor.png" alt="am2320_exporter" width="128" height="128">
</div>

# am2320_exporter  
Prometheus Exporter that acquires temperature and humidity from the sensor [AM2320](http://akizukidenshi.com/download/ds/aosong/AM2320.pdf).  

## Description
This Exporter is created on the assumption that it will run on a Raspberry Pi.  
For temperature and humidity monitoring using Raspberry Pi and Docker.  
It is also possible to deploy on Kubernetes.  

## Exported metrics
| metric                   | description                      | labels |
| ------------------------ | -------------------------------- | ------ |
| am2320_temperature_gauge | temperature obtained from am2320 |        |
| am2320_humidity_gauge    | humidity obtained from am2320    |        |

## Requirements
- Go 1.12.x ~ (Your computer or Raspberry Pi)
- Docker (Raspberry Pi)

## Usage
**The connection between the Raspberry Pi and the AM2320 sensor has been completed in advance.**

### Go
Build the program on your computer or Raspberry Pi.  
`export GO111MODULE=on`  
`GOOS=linux GOARCH=arm GOARM=7 go build`  

Binary execution.  
`sudo ./am2320-exporter`  

### Docker
`docker run -i -t -d --name am2320_exporter -p 9430:9430 --privileged yudaishimanaka/am2320-exporter-armv7l`

### Kubernetes
Configure Pod Affinity according to your Kubernetes cluster and deploy appropriately.  

example
```yml
# Configuration example when deploying to a cluster consisting of only Master and Edge nodes.
apiVersion: extensions/v1beta1
kind: Daemonset
metadata:
  name: am2320-exporter
  namespace: monitoring
  labels:
    name: am2320-exporter
spec:
  template:
    metadata:
      labels:
        app: am2320-exporter
      annotations:
        prometheus.io/scrape: 'true'
        prometheus.io/port: '9430'
        prometheus.io/path: /metrics
    spec:
      containers:
      - name: am2320-exporter
        image: yudaishimanaka/am2320-exporter-armv7l:latest
        imagePullPolicy: Always
        securityContext:
          privileged: true
        ports:
        - containerPort: 9430
      hostNetwork: true
      hostPID: true
```

## License
The MIT License (MIT) -see `LICENSE` for more details.
