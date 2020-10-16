echo 'Starting Envoy front-proxy container'
docker run -d --name front-proxy --network envoymesh -p 9090:90 -p 8001:8001 -e LISTEN_PORT='90' -e SERVICE_DISCOVERY_ADDRESS=service -e SERVICE_DISCOVERY_PORT=90 front-proxy

echo 'Starting Envoy sidecar 1'
docker run -d --name sidecar-1 --network envoymesh --net-alias service -e LISTEN_PORT=90 -e SERVICE_ADDRESS=service-1 -e SERVICE_PORT=9090 sidecar

echo 'Starting Envoy sidecar 2'
docker run -d --name sidecar-2 --network envoymesh --net-alias service -e LISTEN_PORT=90 -e SERVICE_ADDRESS=service-2 -e SERVICE_PORT=9090 sidecar

echo 'Starting service 1'
docker run -d --name service-1 --network envoymesh --net-alias service-1 -e SERVICE_NAME=service-1 service

echo 'Starting service 2'
docker run -d --name service-2 --network envoymesh --net-alias service-2 -e SERVICE_NAME=service-2 service

echo 'Up and running...'

read -n 1 -p "Press any key to exit"