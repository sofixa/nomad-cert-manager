.PHONY: all consul vault nomad

all: consul vault nomad 

consul:
	consul agent -dev -client "0.0.0.0" &
	sleep 5
	
vault:
	vault server -dev &
	sleep 5
	VAULT_ADDR='http://127.0.0.1:8200' vault policy write nomad-server test/nomad-server-policy.hcl
	VAULT_ADDR='http://127.0.0.1:8200' vault policy write nomad-cluster test/nomad-cluster-policy.hcl
	VAULT_ADDR='http://127.0.0.1:8200' vault write /auth/token/roles/nomad-cluster @test/nomad-cluster-role.json

nomad:
	export VAULT_TOKEN="`cat ~/.vault-token`" && nomad agent -dev -consul-address=127.0.0.1:8500 -bind "0.0.0.0" -config test/nomad-options.hcl
	sleep 5


clean:
	pkill consul || true
	pkill vault || true 
	pkill nomad || true
