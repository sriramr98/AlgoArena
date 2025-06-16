prepare:
	npm install
	cd server && go mod tidy
	cd client && npm install