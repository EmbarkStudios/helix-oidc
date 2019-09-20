# Helix Perforce Google OIDC validator
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FEmbarkStudios%2Fhelix-oidc.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FEmbarkStudios%2Fhelix-oidc?ref=badge_shield)
[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)
[![Embark](https://img.shields.io/badge/embark-open%20source-blueviolet.svg)](https://github.com/EmbarkStudios)


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

## Contributing

We welcome community contributions to this project.

Please read our [Contributor Guide](CONTRIBUTING.md) for more information on how to get started.

## License

Licensed under either of

* Apache License, Version 2.0, ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
* MIT license ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

at your option.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FEmbarkStudios%2Fhelix-oidc.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FEmbarkStudios%2Fhelix-oidc?ref=badge_large)

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in the work by you, as defined in the Apache-2.0 license, shall be dual licensed as above, without any additional terms or conditions.
