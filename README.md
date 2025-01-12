# RP42

A Discord Rich Presence integration for [@42School](https://github.com/42School).  
![Screenshot](https://raw.githubusercontent.com/alexandregv/RP42/master/screenshot.png)

## Installation

1. Download RP42 from the [releases](https://github.com/alexandregv/RP42/releases) page, or build it yourself
2. Create an API App on the Intranet: https://profile.intra.42.fr/oauth/applications/new

/!\ Do NOT share your API credentials to someone else, or on GitHub, etc. /!\

## Usage

Run the app like this, using the credentials of your API App:

- Linux: `./RP42 -i CLIENT_ID -s CLIENT_SECRET &`
- Windows: `RP42.exe -i CLIENT_ID -s CLIENT_SECRET`
- MacOS: `open RP42.app -i CLIENT_ID -s CLIENT_SECRET`

## Building yourself

If you want to build RP42 yourself, follow these instructions:

1. Clone the repo: `git clone https://github.com/alexandregv/RP42.git`
2. Compile: `make` for all distro, or `make linux` / `make macos` / `make windows`

## Contributing

1. Fork it (<https://github.com/alexandregv/RP42/fork>)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## Contributors

- [alexandregv/aguiot--](https://github.com/alexandregv) - creator and maintainer

## Stargazers over time

[![Stargazers over time](https://starchart.cc/alexandregv/RP42.svg?variant=adaptive)](https://starchart.cc/alexandregv/RP42)
