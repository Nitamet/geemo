<div align="center" style="text-align: center">
  <img src="https://geemo.app/appicon.png" height="50%" width="50%"/>
  <h3><strong>geemo</strong></h3>
  A helper app to import builds from popular sites to League of Legends app.
  Written in <strong>Go</strong> & <strong>TypeScript</strong>, using <a href="https://github.com/wailsapp/wails">Wails</a> and <a href="https://vuejs.org/">Vue 3</a>.
  <br/>
  <br/>
  <img src="https://geemo.app/app.png"/>
</div>

## Note 
This project is in its early stages of its existence, so expect some bugs. Please report any errors you find to the issues page if possible.

## Features
- Import builds directly into your LoL client on the fly
- Multiple data sources to choose builds from
  - __u.gg__
  - __Mobalytics__
  - ...more sources will be added in the future
- __Summoner's Rift / ARAM__ support
- Different roles have different builds
- Role is detected automatically in Draft Pick mode
- Windows & __Linux__ support (Mac is not supported yet but it shouldn't be hard to add)
- Relatively lightweight: binary size is ~9 MB; RAM usage is 130-160 MB
- No ads & no tracking
- Self-updating
- i18n support
- - English
  - Russian
  - ...

 ## Download
 __Note:__ On Windows, [webview2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download-section) is required to run this app
 - Download the latest [release](https://github.com/Nitamet/geemo/releases/latest) for your OS.
 - Unzip the archive.

## How to use
 - Run geemo
 - Run the LoL client
 - Get into a game lobby and select a champion
 - Pick a build and import it

## Development
geemo is using [Wails](https://github.com/wailsapp/wails) framework, which allows to write desktop applicatios in Go and JavaScript.

### Requirements:
- Go
- NPM
- Wails CLI

Then, in the project's root directory, you can run this command
```wails dev```
which will run the application in live development mode.

Alternatively, you can run ```wails build``` command which will compile the binary into the ```build/bin``` folder.
