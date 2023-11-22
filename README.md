# SIEM project

## Technologies
### VM1 - ELK Server

- Docker				: ?
- ELK(Elasticsearch, Logstash, Kibana)	: 7.17.15

### VM2 - Weak Server

- Filebeat				: 7.17.15
- Snort					: 2.9.7.0-5build1
- Nginx					: ?


## Running the project
### VM1 - ELK Server 
VM that will install the Docker with Elasticsearch, Logstash and Kibana containers.  

Specs: VMBox Lubuntu with +-6GB Ram, 2 processors and ~15GB storage
Configure the Bridged Adapter in Virtual Box to stay on the same host network

1. Install docker
https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository

2. Clone the project
$ git clone https://github.com/Reverendyz/elk.git

3. Start the ELK
$ cd elk && docker compose up

4. Check Elasticsearch and Kibana
$ curl -s -I http://localhost:9200 | grep 'HTTP/1.1 200 OK'
$ curl -s -I http://localhost:5601 | grep 'HTTP/1.1 302 Found' 

You can open Kibana in browser too - http://X.X.X.X:5601
X.X.X.X is the IP Address of this VM


### VM2 - WeakServer
VM with Snort and Filebeat do detect and send alerts to ELKServer respectively and with Nginx server.

Specs: VMBox Lubuntu with +-4GB Ram, 2 processors and ~10GB storage
Configure the Bridged Adapter to stay on the same host network and Promiscuous Mode to "Allow all"

1. Basic dependencies
$ apt install openssh-server curl net-tools vim openjdk-8-jdk nginx snort apt-transport-https -y
$ java -version

2. Snort
More information about configuration read these links 
- https://www.securityarchitecture.com/learning/intrusion-detection-systems-learning-with-snort/configuring-snort-on-linux/ 
- https://www.youtube.com/watch?v=Osu_PEA94mI 
- https://alparslanakyildiz.medium.com/writing-custom-snort-rules-e9abe10932e1 

+ Edit the snort.conf with these configurations
$ vim /etc/snort/snort.conf
  ipvar HOME_NET X.X.X.X 
  ipvar EXTERNAL_NET !$HOME_NET 
  ...
  config logdir: /var/log/snort


X.X.X.X means the host or the network we want to observe and !$HOME_NET means all the external 
network will be any IP address different of that or isn’t part of home network 

+ (Optional )Download via browser the snort rules from snort website and extract then to snort rules folder 
$ tar –xvzf community_rules.tar.gz -C /etc/snort/rules 

+ Set promiscuous mode in our network interface 
$ ip link set enp0s3 promisc on 
 
+ Open the /etc/snort/rules/local.rules, erase everyhing and add our rules 
\# 1. Requisição de uma página web 
alert tcp any any -> X.X.X.X 80 (msg: "Requisição página web"; sid:1; content:"|47 45 54|"; offset:0; depth:3;) 
 
\# 2. Fim de uma conexão no SSH 
alert tcp any any -> X.X.X.X 22 (msg: "Fim de conexão SSH"; sid:2; flags:FA;) 
 
X.X.X.X means the IP of this host 
 
+ Test the Snort configuration using the network interface(in my case is enp0s3). Basically, we are telling Snort to test (-T) the following rules file (-c) while it is listening on the enp0s3 interface (-i).

$ snort -T -c /etc/snort/snort.conf -i enp0s3 

Check if results are ok 


3. Filebeat
Filebeat is a lightweight log shipper (expedidor, que envia), which will reside on the same instance as the Apache, Nginx or Snort. 

+ Download and import the GPG public key from Elasticsearch to APT 
$ wget -prefer-family=IPv4 -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -

+ Add the Elastic origin list to the dir sources.list.d, where the APT will search for new origins 
$ sudo sh -c 'echo "deb https://artifacts.elastic.co/packages/7.x/apt stable main" > /etc/apt/sources.list.d/elastic-7.x.list' 

+ Update and install
$ apt update 
$ apt install filebeat 

+ Open the /etc/filebeat/filebeat.yml and comment this part with ‘#’ 
\#output.elasticsearch: 
    # Array of hosts to connect to. 
    # hosts: ["localhost:9200"] 

Uncomment this another part by removing ‘#’ 
output.logstash: 
\# The Logstash hosts 
hosts: ["X.X.X.X:5044"] 
 
Use the ELKServer IP Address instead of X.X.X.X.  
 
Change this part too like this 
filebeat.inputs: 
\# Each - is an input. Most options can be set at the input level, so 
\# you can use different inputs for various configurations. 
\# Below are the input specific configurations. 
 
\# filestream is an input for collecting log messages from files. 
\- type: log 
 
  # Unique ID among all inputs, an ID is required. 
  id: my-filestream-id 
 
  # Change to true to enable this input configuration. 
  enabled: true 
 
  # Paths that should be crawled and fetched. Glob based paths. 
  paths: 
    - /var/log/snort/alert 
 
+ Save, close file and run these commands below 
$ filebeat modules enable snort 
$ filebeat setup --pipelines --modules snort 
$ filebeat setup --index-management -E output.logstash.enabled=false -E 'output.elasticsearch.hosts=["X.X.X.X:9200"]' 
$ filebeat setup -E output.logstash.enabled=false -E output.elasticsearch.hosts=['X.X.X.X:9200'] -E setup.kibana.host= X.X.X.X:5601 
 
$ systemctl start filebeat 
$ systemctl enable filebeat 

+ Start Snort to sniff our network 
$ snort -A fast -c /etc/snort/snort.conf -q -i enp0s3 
 
\-A:    Alert  using  the  specified  alert-mode.  Valid alert modes include fast, full, none, and unsock.  Fast writes  alerts  to  the  default "alert"  file  in  a  single-line, syslog style alert message.  Full writes the alert to the "alert" file with the full decoded header as well  as  the alert message.  None turns off alerting.  Unsock is an experimental mode that sends the alert information out over  a  UNIX socket to another process that attaches to that socket. 

\-q:  Quiet  operation.   Don't display banner and initialization information. 


### Testing
+ In ELKServer check if data is comming from Logstash to Elasticsearch
$ curl -XGET 'http://localhost:9200/filebeat-\*/\_search?pretty' | less

Look for "hits" in JSON response

+ Open kibana in browser X.X.X.X:5601, open left menu and click in “Discovery” item 
Select filebeat-\* in filter, change the time filter to “Last 30 minutes” and check the logs on chart
