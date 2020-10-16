echo "Stopping containers"
docker rm -f front-proxy sidecar-1 sidecar-2 service-1 service-2

echo "Stopped"

read -n 1 -p "Press any key to exit"