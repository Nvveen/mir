all: sprinter

sprinter:
	cd cmd/sprinter; mkdir -pv sprinter/sprinter
	mkdir -pv cmd/sprinter/sprinter/sprinter cmd/sprinter/sprinter/storage/
	cd cmd/sprinter/sprinter/sprinter; go build ../../
	cp -R storage/mongo_test/ cmd/sprinter/sprinter/storage

archive: sprinter
	cd cmd/sprinter; tar czvf ../../sprinter.tar.gz sprinter/

sink:
	cd cmd/sink; go build .

clean:
	rm -rf cmd/sprinter/sprinter
	rm cmd/sink/sink

.PHONY: sprinter
