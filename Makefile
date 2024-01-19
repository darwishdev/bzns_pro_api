# Include the main .env file
include config/state.env
# Construct the variable name based on STATE
CURRENT_STATE_FILE = config/$(STATE).env
# Include the appropriate .env file (e.g., dev.env or prod.env)
include $(CURRENT_STATE_FILE)

# Include the additional .env file
include config/shared.env
 

test:
	go test  -v -race  -cover ./... 

seed: 
<<<<<<< HEAD
	curl -X POST ${SEEDER_URL}/storage/storage && curl -X POST ${SEEDER_URL}/accounts/settings && curl -X POST ${SEEDER_URL}/accounts/permissions && curl -X POST ${SEEDER_URL}/accounts/roles && curl -X POST ${SEEDER_URL}/accounts/users && curl -X POST ${SEEDER_URL}/accounts/navigations   
 
=======
	    curl -X POST ${SEEDER_URL}/places  &&  curl -X POST ${SEEDER_URL}/storage  && curl -X POST ${SEEDER_URL}/entities/entities && curl -X POST ${SEEDER_URL}/entities/halls && curl -X POST ${SEEDER_URL}/accounts/permissions && curl -X POST ${SEEDER_URL}/accounts/roles && curl -X POST ${SEEDER_URL}/accounts/accounts && curl -X POST ${SEEDER_URL}/accounts/users && curl -X POST ${SEEDER_URL}/accounts/navigations && curl -X POST ${SEEDER_URL}/products/units && curl -X POST ${SEEDER_URL}/products/categories && curl -X POST ${SEEDER_URL}/products/products && curl -X POST ${SEEDER_URL}/products/ingredients && curl -X POST ${SEEDER_URL}/products/modifiers && curl -X POST ${SEEDER_URL}/entities/settings && curl -X POST ${SEEDER_URL}/devices/devices && curl -X POST ${SEEDER_URL}/devices/settings  && curl -X POST ${SEEDER_URL}/entities/settings

>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
seed_t:
	curl -X POST http://192.168.1.40:3000/users -d '{"test":true}' -H "Content-Type: application/json"
	
db_s:
	supabase start -x realtime 
mign : 
	supabase migration new $(name)
migu : 
	supabase migration up 
migd : 
	migrate -path common/db/migration -database ${DB_SOURCE} -verbose down 

cdb: 
	docker exec -it postgres  createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}
	
ddb :
	docker exec -it postgres  dropdb --username=${DB_USER}   ${DB_NAME}  --force
rdb:
	supabase db reset
rdbr:
	supabase db reset --linked

rdbs:
	make rdb seed
rdb_t:
	make  ddb cdb migu seed_t
run:
	go run main.go



mock:
	mockgen -package mockdb -destination common/db/mock/store.go github.com/meloneg/mln_rms_core/common/db/gen Store
sqlc :
	rm -rf common/db/gen/*.sql.go && sqlc generate


buf:
	rm -rf common/pb/* && buf generate



buf_push_g:
	cd common/proto && git add . && git commit -m "sync" && git push -u origin main
buf_push:
	cd common/proto && buf push



buf_u:
	buf mod update


dtag:
	docker tag mln_api_core exploremelon/mln_api_core:${v}

dpush:
	docker push  exploremelon/mln_api_core:${v}
 

<<<<<<< HEAD
	docker push exploremelon/mln_api_core:

=======
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
run_bg:
	make run  > /dev/null 2>&1 & && disown


gen:
<<<<<<< HEAD
	rm -rf lib/pb/*  && protoc -I=proto/mln_api_protos --dart_out=grpc:lib/pb proto/mln_api_protos/bznspro/v1/*.proto proto/mln_api_protos/bznspro/v1/*/*.proto  proto/mln_api_protos/google/protobuf/*.proto
=======
	rm -rf lib/pb/*  && protoc -I=proto/mln_api_protos --dart_out=grpc:lib/pb proto/mln_api_protos/rms/v1/*.proto proto/mln_api_protos/rms/v1/*/*.proto  proto/mln_api_protos/google/protobuf/*.proto
>>>>>>> 11dce109f0ac477a16b39aab62601d26ece07212
