<div align="center" style="text-align: center">
  <img src="https://geemo.app/appicon.png" height="30%" width="30%"/>
  <h3><strong>geemo</strong></h3>
  <div>
    <a href="https://github.com/Nitamet/geemo/releases/latest"><img alt="GitHub release (with filter)" src="https://img.shields.io/github/v/release/Nitamet/geemo?style=for-the-badge&logo=github"></a>
    <a href="https://github.com/Nitamet/geemo/actions/workflows/release-app.yaml"><img alt="Build" src="https://img.shields.io/github/actions/workflow/status/Nitamet/geemo/release-app.yaml?style=for-the-badge&logo=github"/></a>
    <a href="https://github.com/Nitamet/geemo/blob/main/LICENSE"><img alt="GitHub" src="https://img.shields.io/github/license/Nitamet/geemo?style=for-the-badge&logo=github"></a>
    <a href="https://goreportcard.com/badge/github.com/Nitamet/geemo"><img alt="Go Report" src="https://goreportcard.com/badge/github.com/Nitamet/geemo?style=for-the-badge"></a>
  </div>
  A helper app to import builds from popular sites to League of Legends app.
  <br/>
  Written in <strong>Go</strong> & <strong>TypeScript</strong>, using <a href="https://github.com/wailsapp/wails">Wails</a> and <a href="https://vuejs.org/">Vue 3</a>.
  <br/>
  <br/>
  <img src="https://geemo.app/app.png"/>
</div>

## Note 
This project is in its early stages of existence, so expect some bugs. Please report any errors you find to the issues page if possible.

## Features
- Import builds directly into your LoL client on the fly
- Multiple data sources to choose builds from
  - __u.gg__
  - __Mobalytics__
  - ...more sources will be added in the future
- __Summoner's Rift / ARAM__ support
- Different roles have different builds
- Role is automatically detected in Draft Pick mode
- Windows & __Linux__ support (Mac is not supported yet but it shouldn't be hard to add)
- Relatively lightweight: binary size is ~9 MB; RAM usage is 130-160 MB
- No ads & no tracking
- Self-updating
- i18n support
  - English
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

## Disclaimer
geemo isn’t endorsed by Riot Games and doesn’t reflect the views or opinions of Riot Games or anyone officially involved in producing or managing League of Legends. League of Legends and Riot Games are trademarks or registered trademarks of Riot Games, Inc. League of Legends © Riot Games, Inc.
