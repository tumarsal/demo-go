rem docker kill $(docker ps -q)
rem docker rm $(docker ps -a -q)
docker rm demo-app-run
docker rmi demo-golang-app
docker build -t demo-golang-app .
docker run --publish 6060:8000 --rm --name demo-app-run demo-golang-app
explorer "http://localhost:6060"
pause