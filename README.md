# DLI

```
A CLI tool that allows you to manage DNS records across various DNS providers using environment variables for authentication.

Usage:
  dli [command]

Available Commands:
  append      Append a DNS record
  completion  Generate the autocompletion script for the specified shell
  delete      Delete a DNS record
  help        Help about any command
  list        List DNS records
  list-zones  List DNS zones
  set         Set a DNS record

Flags:
  -h, --help              help for dli
      --provider string   DNS provider (e.g., bunny, cloudflare, route53)
      --zone string       DNS zone to manage

Use "dli [command] --help" for more information about a command.
```

## All LibDNS Providers

Below are all the LibDNS providers.
Many are not yet updated to support libdns 1.1.0 so are not implemented in this project yet.
Environment variable names are chosen to match https://github.com/go-acme/lego/ to allow the easy reuse of environment files.

- [!] **alidns** (incompatible with current libdns - Record struct interface mismatch)
- [!] **autodns** (incompatible with current libdns - Record struct interface mismatch)
- [x] **azure** - Environment variables: `AZURE_SUBSCRIPTION_ID` + `AZURE_RESOURCE_GROUP` + optional `AZURE_TENANT_ID` + optional `AZURE_CLIENT_ID` + optional `AZURE_CLIENT_SECRET`
- [x] **bunny** - Environment variable: `BUNNY_API_KEY`
- [!] **civo** (incompatible with current libdns - Record struct interface mismatch)
- [x] **cloudflare** - Environment variable: `CF_DNS_API_TOKEN` or `CLOUDFLARE_DNS_API_TOKEN`
- [x] **cloudns** - Environment variables: `CLOUDNS_AUTH_ID` + `CLOUDNS_AUTH_PASSWORD` + optional `CLOUDNS_SUB_AUTH_ID`
- [!] **ddnss** (incompatible with current libdns - Record struct interface mismatch)
- [x] **desec** - Environment variable: `DESEC_TOKEN`
- [x] **digitalocean** - Environment variable: `DO_AUTH_TOKEN`
- [!] **dinahosting** (incompatible with current libdns - Record struct interface mismatch)
- [!] **directadmin** (incompatible with current libdns - Record struct interface mismatch)
- [!] **dnsexit** (incompatible with current libdns - Record struct interface mismatch)
- [!] **dnsmadeeasy** (incompatible with current libdns - Record struct interface mismatch)
- [!] **dnspod** (incompatible with current libdns - Record struct interface mismatch)
- [x] **dnsimple** - Environment variable: `DNSIMPLE_OAUTH_TOKEN` + optional `DNSIMPLE_BASE_URL`
- [!] **dode** (incompatible with current libdns - only supports RecordAppender and RecordDeleter)
- [!] **dreamhost** (incompatible with current libdns - Record struct interface mismatch)
- [x] **duckdns** - Environment variable: `DUCKDNS_TOKEN`
- [!] **dynv6** (incompatible with current libdns - Record struct interface mismatch)
- [x] **dynu** - Environment variable: `DYNU_API_KEY`
- [!] **easydns** (incompatible with current libdns - Record struct interface mismatch)
- [!] **exoscale** (incompatible with current libdns - Record struct interface mismatch)
- [ ] **gandi**
- [ ] **gcore**
- [x] **glesys** - Environment variables: `GLESYS_API_USER` + `GLESYS_API_KEY`
- [ ] **godaddy**
- [x] **googleclouddns** - Environment variables: `GCE_PROJECT` + `GCE_SERVICE_ACCOUNT` or `GCE_SERVICE_ACCOUNT_FILE`
- [x] **he** - Environment variable: `HURRICANE_TOKENS` or `HE_API_KEY`
- [x] **hetzner** - Environment variable: `HETZNER_API_KEY`
- [ ] **hexonet**
- [!] **hosttech** (incompatible with current libdns - Record struct interface mismatch)
- [!] **huaweicloud** (incompatible with current libdns)
- [x] **inwx** - Environment variables: `INWX_USERNAME` + `INWX_PASSWORD` + optional `INWX_SHARED_SECRET`
- [ ] **ionos**
- [ ] **leaseweb**
- [ ] **linode**
- [x] **loopia** - Environment variables: `LOOPIA_API_USER` + `LOOPIA_API_PASSWORD`
- [x] **luadns** - Environment variables: `LUADNS_API_USERNAME` + `LUADNS_API_TOKEN`
- [x] **mailinabox** - Environment variables: `MAILINABOX_BASE_URL` + `MAILINABOX_EMAIL` + `MAILINABOX_PASSWORD`
- [x] **metaname** - Environment variables: `METANAME_API_KEY` + `METANAME_ACCOUNT_REFERENCE`
- [x] **mijnhost** - Environment variable: `MIJNHOST_API_KEY`
- [!] **mythicbeasts** (incompatible with current libdns - Record struct interface mismatch)
- [!] **namecheap** (incompatible with current libdns)
- [ ] **namedotcom**
- [ ] **namesilo**
- [x] **netcup** - Environment variables: `NETCUP_CUSTOMER_NUMBER` + `NETCUP_API_KEY` + `NETCUP_API_PASSWORD`
- [ ] **netlify**
- [x] **nfsn** - Environment variables: `NEARLYFREESPEECH_LOGIN` + `NEARLYFREESPEECH_API_KEY`
- [ ] **nicrudns**
- [ ] **njalla**
- [ ] **openstack-designate**
- [x] **ovh** - Environment variables: `OVH_APPLICATION_KEY` + `OVH_APPLICATION_SECRET` + `OVH_CONSUMER_KEY` + `OVH_ENDPOINT`
- [x] **porkbun** - Environment variables: `PORKBUN_API_KEY` + `PORKBUN_SECRET_API_KEY`
- [ ] **powerdns**
- [ ] **rfc2136**
- [!] **route53** (incompatible with current libdns)
- [x] **scaleway** - Environment variables: `SCW_SECRET_KEY` + `SCW_PROJECT_ID`
- [ ] **selectel**
- [ ] **tencentcloud**
- [ ] **timeweb**
- [ ] **totaluptime**
- [ ] **transip**
- [ ] **vercel**
- [x] **vultr** - Environment variable: VULTR_API_KEY
- [x] **westcn** - Environment variables: `WESTCN_USERNAME` + `WESTCN_API_PASSWORD`
