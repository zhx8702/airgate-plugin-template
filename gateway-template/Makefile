.PHONY: build build-web build-backend manifest clean

build: build-web build-backend

build-web:
	cd ../airgate-sdk/frontend && npm run build
	cd web && npm run build

build-backend: build-web
	rm -rf backend/internal/gateway/webdist
	cp -r web/dist backend/internal/gateway/webdist
	cd backend && go build -o ../bin/plugin-template .

manifest:
	cd backend && go run ./cmd/genmanifest/

clean:
	rm -rf bin/ web/dist/ backend/internal/gateway/webdist/
