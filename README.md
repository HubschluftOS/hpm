## Hubshluft Package Manager

Lightweight and powerful package manager for Hubshluft and other GNU/Linux distributions

## Installation

1. **The first way.** Download the latest release.

```bash
curl -LO https://github.com/hubshluft/hpm/releases/download/v?.?.?/hpm-?.?.?.tar.gz
tar xzf hpm-*.tar.gz
```

2. **The second way.** Use `make` to build `hpm`. [go, make]

```bash
git clone https://github.com/hubshluft/hpm
cd hpm
make
```


## Dependencies

`sudo, sh`

## Usage

```bash
get	 Install the needed package on your Linux system
unget	 Uninstall the package on your Linux system
update	 Update all packages on your Linux system
news	 Display the latest news of the Hubshluft team and HubshluftOS
help	 Display extra information about package manager
```

## License

[BSD 2-Clause License](https://opensource.org/license/bsd-2-clause)
