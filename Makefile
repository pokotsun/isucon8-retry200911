branch=master
APP_HOME := $(HOME)/torb
build: 
	(cd $(APP_HOME) && git reset --hard HEAD && git fetch && git checkout ${branch} && git pull origin ${branch})
	(cd $(APP_HOME)/webapp/go && make build)
	sudo systemctl restart torb.go 
analyze:
	# sudo cp /dev/null /var/log/mysql/mysql-slow.log
	# sudo cp /dev/null /var/log/nginx/access.log
	(cd $(APP_HOME)/bench && ./bin/bench && jq . < result.json)
	# sudo alp --file=/var/log/nginx/access.log ltsv -r --sort sum -m "/diary/entries/.+"| head -n 30
.PHONY: build analyze

