# Helix Perforce Google OIDC validator
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FEmbarkStudios%2Fhelix-oidc.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FEmbarkStudios%2Fhelix-oidc?ref=badge_shield)


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


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FEmbarkStudios%2Fhelix-oidc.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FEmbarkStudios%2Fhelix-oidc?ref=badge_large)