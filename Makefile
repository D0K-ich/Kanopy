GO ?= go

ANSIBLE_PARAMS=-D -f 3
STAGE_PATH=./deploy/inv/stage

BENCH_FLAGS = -v -cpu=1 -check.v -check.b -check.bmem -check.btime=2s ${EXTRA}

RL_TAG=`git describe --tags $(git rev-list --tags --max-count=1)`
COM_ID=`git log --format="%h" -n 1`

prod-build:
	@echo "===================================="
	@echo "Building main app..."
	@echo "------------------------------------"
	cd /Users/yjrur/IdeaProjects/Kanopy/view/ & wails build
	cd /Users/yjrur/IdeaProjects/Kanopy/view/build/bin/ & .\kanopy.exe

prod-build:
	@echo "===================================="
	@echo "Building main app..."
	@echo "------------------------------------"
	cd /Users/yjrur/IdeaProjects/Kanopy/view/ & wails build
	cd /Users/yjrur/IdeaProjects/Kanopy/view/build/bin/ & .\kanopy.exe



#help:					# display this help
#	@echo "Targets"
#	@echo "-------"
#	@grep -iE '^[a-z\-]+:' Makefile | sed -Ee 's/://; s/#/-/'

#prod-setup:
#	@echo "===================================="
#	@echo "Setup manager backend infrastructure"
#	@echo "------------------------------------"
#	ansible-playbook ${ANSIBLE_PARAMS} -i ${STAGE_PATH} ./deploy/pb/setup.backend.prod.yml --limit=manager-prod
#
#prod-build:
#	@echo "============================"
#	@echo "Building manager PROD binary"
#	@echo "----------------------------"
#	go build -v -o bin/manager.prod ./cmd/service
#
#prod-deploy: test-fetch-rules
#	@echo "====================================="
#	@echo "Deploy manager backend infrastructure"
#	@echo "-------------------------------------"
#	ansible-playbook ${ANSIBLE_PARAMS} -i ${STAGE_PATH} ./deploy/pb/deploy.backend.prod.yml --limit=manager-prod
#
#dev-setup:
#	@echo "========================================"
#	@echo "Setup manager dev backend infrastructure"
#	@echo "----------------------------------------"
#	ansible-playbook ${ANSIBLE_PARAMS} -i ${STAGE_PATH} ./deploy/pb/setup.backend.dev.yml --limit=manager-dev
#
#dev-build:
#	@echo "==========================="
#	@echo "Building manager DEV binary"
#	@echo "---------------------------"
#	go build -v -o bin/manager.dev ./cmd/service
#
#dev-deploy: test-fetch-rules
#	@echo "========================================="
#	@echo "Deploy manager dev backend infrastructure"
#	@echo "-----------------------------------------"
#	ansible-playbook ${ANSIBLE_PARAMS} -i ${STAGE_PATH} ./deploy/pb/deploy.backend.dev.yml --limit=manager-dev
#
#count-code:
#	@echo "========================"
#	@echo "Gathering code statistic"
#	git ls-files | xargs cloc
#
#test-fetch-rules:
#	go test -v ./frules_test.go

#run-yarik:
#	@echo "==========================="
#	@echo "Running manager PROD binary"
#	@echo "---------------------------"
#	./bin/manager.prod --config=templates/manager.yarik.yml
#
#run-danik:
#	@echo "==========================="
#	@echo "Running manager PROD binary"
#	@echo "---------------------------"
#	./bin/manager.prod --config=templates/manager.danik.yml

#prod-run:
#	@echo "============================="
#	@echo "Running itemprice PROD binary"
#	@echo "-----------------------------"
#	./bin/manager.prod

#front-get:
#	@echo "============================="
#	@echo "Gettings itemprice FRONT repo"
#	@echo "-----------------------------"
#	git clone https://gitbit.ru/ncrawler/ip-front front
#
#front-deploy:
#	@echo "============"
#	@echo "DEPLOY FRONT"
#	@echo "------------"
#	cd front && npm run deploy


#extract:
#	tar -xvf ./gochrome-119.tgz -C ./bin/
#
#
#test-detection:
#	go test -v -run "TestDectection"
#
#test-worker:
#	go test -v -run "TestWorker"