docker-pn:
	sudo docker-compose stop primary-node && sudo docker-compose rm -f primary-node
	sudo docker-compose up --build primary-node

docker-dn1:
	sudo docker-compose stop data-node-1 && sudo docker-compose rm -f data-node-1
	sudo docker-compose up --build data-node-1

docker-dn2:
	sudo docker-compose stop data-node-2 && sudo docker-compose rm -f data-node-2
	sudo docker-compose up --build data-node-2

docker-diaboromon:
	sudo docker-compose stop diaboromon && sudo docker-compose rm -f diaboromon
	sudo docker-compose up --build diaboromon

docker-tai:
	sudo docker-compose stop tai && sudo docker-compose rm -f tai
	sudo docker-compose up --build -d tai
	sudo docker attach tai_container

docker-if:
	sudo docker-compose stop isla-file && sudo docker-compose rm -f isla-file
	sudo docker-compose up --build isla-file

docker-cf:
	sudo docker-compose stop continente-folder && sudo docker-compose rm -f continente-folder
	sudo docker-compose up --build continente-folder

docker-cs:
	sudo docker-compose stop continente-server && sudo docker-compose rm -f continente-server
	sudo docker-compose up --build continente-server