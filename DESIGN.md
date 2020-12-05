# Worflow
* listen for new Deployments
* parse new Deployments' tags to check for new domains that need certificates and regularly poll Consul services
* update Consul with configuration (new domains, accounts, challenge to use, etc.)
* demand new certificate
* respond to challenge
    * DNS can be done directly
    * HTTP and tls-alpn-01 we'll see
* write result to Vault

# Components
* Nomad's [https://github.com/hashicorp/nomad/tree/master/api](API package) for [https://www.hashicorp.com/blog/building-on-top-of-hashicorp-nomad-s-event-stream](Nomad Event stream) to listen for new Deployments
* Consul's [https://github.com/hashicorp/consul/tree/master/api](API package) for interacting with Consul (for configuration and services)
* an ACME library for the ACME part
* Vault's [https://github.com/hashicorp/vault/tree/master/api](API package) for integrating with Vault (storing certificates)

