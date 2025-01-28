.DEFAULT_GOAL = run

create-vm-win:
	pwsh scripts/create-vm.ps1

destroy-vm-win:
	pwsh scripts/destroy-vm.ps1

create-vm:
	VAGRANT_CWD=$(CURDIR)/test/vagrant vagrant up
	VAGRANT_CWD=$(CURDIR)/test/vagrant vagrant ssh -c "sudo reboot"

destroy-vm:
	VAGRANT_CWD=$(CURDIR)/test/vagrant vagrant destroy -f
	
create-venv:
	$(CURDIR)/scripts/create-venv.sh

run:
	make swg && \
	go run cmd/main.go

swg:
	swag init -g ./cmd/main.go -o docs