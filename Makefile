kxsystems:

clean:
	docker-compose down
	docker rmi gateway-service-img storage-service-img-1 storage-service-img-2 storage-service-img-3

main:
	cd gatewayservice && rm -f main && make
	cd storageservice && rm -f main && make

run:
	docker-compose up

