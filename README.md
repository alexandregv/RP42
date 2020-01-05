# RP42

A Discord Rich Presence integration for [@42School](https://github.com/42School).  
![Screenshot](https://raw.githubusercontent.com/alexandregv/RP42/master/screenshot.png)

## Installation
If you are logged at 42Paris or 42Lyon, you don't have to download or install anything. Just skip to the Usage instructions.  
Otherwise, download RP42 from the [releases](https://github.com/alexandregv/RP42/releases) page, or build it yourself.  

## Usage at 42Paris or 42Lyon
Paris: execute `open /sgoinfre/goinfre/Perso/aguiot--/public/RP42.app`  
Lyon: execute `open /sgoinfre/Perso/amonteli/RP42/RP42.app`.  
The app will appear in your menu/tray bar (near the clock). Right click it to quit.

## Usage in other campuses
Follow Installation instructions and then double-click the file, or use these commands:  
Linux: `./RP42 &`  
Windows: `RP42.exe`  
MacOS: `open RP42.app`  
The app will appear in your menu/tray bar (near the clock). Right click it to quit.

## Building yourself
If you want to build RP42 yourself, you will have to generate an app on the 42's API and then follow these instructions:  
1. Clone the repo: `git clone https://github.com/alexandregv/RP42.git`  
2. Export API credentials: `export RP42_CLIENT_ID=<your-client-id> && export RP42_CLIENT_SECRET=<your-client-secret>`  
3. Compile: `make`  

## Contributing

1. Fork it (<https://github.com/alexandregv/RP42/fork>)  
2. Create your feature branch (`git checkout -b my-new-feature`)  
3. Commit your changes (`git commit -am 'Add some feature'`)  
4. Push to the branch (`git push origin my-new-feature`)  
5. Create a new Pull Request  

## Contributors

- [alexandregv/aguiot--](https://github.com/alexandregv) - creator and maintainer  
