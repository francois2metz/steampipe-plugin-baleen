# Baleen plugin for Steampipe

Use SQL to query namespaces, rules and more from [Baleen][].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install francois2metz/baleen

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/baleen.spc ~/.steampipe/config/baleen.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[baleen]: https://baleen.cloud/
