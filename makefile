compose:
	docker-compose -f docker/docker-compose.yml up -d 

pip:
	pip install mysql-connector-python sqlalchemy