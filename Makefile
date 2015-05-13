all: sprinter

sprinter:
	cd cmd/sprinter; mkdir -pv sprinter/sprinter
	mkdir -pv cmd/sprinter/sprinter/sprinter cmd/sprinter/sprinter/storage/
	cd cmd/sprinter/sprinter/sprinter; go get ../../ && go build ../../
	cp -R storage/mongo_test/ cmd/sprinter/sprinter/storage

archive: sprinter
	cd cmd/sprinter; tar czvf ../../sprinter.tar.gz sprinter/

sink:
	cd cmd/sink; go build .

deliverable: progress report archive
	mkdir -v Neal-vanVeen-s0718971-deliverable
	cp project-progress/project-progress.tex Neal-vanVeen-s0718971-deliverable/
	cp project-progress/project-progress.pdf Neal-vanVeen-s0718971-deliverable/
	cp project-report/project-report.tex Neal-vanVeen-s0718971-deliverable/
	cp project-report/project-report.pdf Neal-vanVeen-s0718971-deliverable/
	cd Neal-vanVeen-s0718971-deliverable; git clone https://github.com/Nvveen/mir
	cp README-deliverable Neal-vanVeen-s0718971-deliverable/README.md
	cp sprinter.tar.gz Neal-vanVeen-s0718971-deliverable/
	tar czvf Neal-vanVeen-s0718971-deliverable.tar.gz Neal-vanVeen-s0718971-deliverable

progress:
	cd project-progress; make && make && make

report:
	cd project-report; make && make && make

clean:
	rm -rf cmd/sprinter/sprinter
	rm -rf cmd/sink/sink
	rm -rf Neal-vanVeen-s0718971-deliverable

.PHONY: sprinter
