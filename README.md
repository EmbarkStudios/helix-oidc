# Helix Perforce Google OIDC validator

## Server Installation
``` bash
$ p4 triggers -o
Triggers:
        oidc-sso auth-check-sso auth "/usr/bin/helix-oidc --google-client-id <google-client-id> validate --id %email% -c %clientip%"
```

## Client Setup
Example using [shelmangroup's oidc-agent](https://github.com/shelmangroup/oidc-agent)
```
p4 set P4LOGINSSO="oidc-agent get -n p4 -o id_token"
p4 login -a
```
