docker-broker:
	sudo docker-compose stop broker && sudo docker-compose rm -f broker
	sudo docker-compose up --build broker

docker-h1:
	sudo docker-compose stop hextech1 && sudo docker-compose rm -f hextech1
	sudo docker-compose up --build hextech1

docker-h2:
	sudo docker-compose stop hextech2 && sudo docker-compose rm -f hextech2
	sudo docker-compose up --build hextech2

docker-h3:
	sudo docker-compose stop hextech3 && sudo docker-compose rm -f hextech3
	sudo docker-compose up --build hextech3

docker-jayce:
	sudo docker-compose stop jayce && sudo docker-compose rm -f jayce
	sudo docker-compose up --build -d jayce
	sudo docker attach jayce_container

docker-s1:
	sudo docker-compose stop supervisor1 && sudo docker-compose rm -f supervisor1
	sudo docker-compose up --build -d supervisor1
	sudo docker attach supervisor1_container

docker-s2:
	sudo docker-compose stop supervisor2 && sudo docker-compose rm -f supervisor2
	sudo docker-compose up --build -d supervisor2
	sudo docker attach supervisor2_container

