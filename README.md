# labs

All labs are preloaded onto the NUCs in front of the room. You will find cards located on desks that indicate the IP address, username and password to perform remote development with VSCode. These machines have all required dependencies to perform the labs. 

> In case you wish to not leverage the machines provided today. You may download the labs from the following repo: 
```
git clone https://github.com/rsdmike/labs.git
```
You must also use the customized docker-compose in the nightly build here: https://github.com/rsdmike/developer-scripts.git on the labs branch
```
git clone https://github.com/rsdmike/developer-scripts.git
git checkout labs
```
To bring up EdgeX use:
`cd ~/Documents/developer-scripts/releases/nightly-build/compose-files && docker-compose -f docker-compose-nexus-no-secty.yml up -d`



## Prerequisites

- Must have VSCode installed
- Must have an SSH installed (typically included with a git installation on Windows)
- Install Remote Development Extension https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack
(You may ignore any WSL errors if you are on windows)

## Connecting Remotely

 1. After you have installed the above extension, add an SSH target with the IP listed on a card.
 ```
 ssh user@[ipaddress]
 ```
 Password: `EdgeX`

 2. Open up `~/Documents/labs`. 
 3. Open up a terminal in VSCode` (Ctrl+~)` -- this terminal is your remote session. You should see something similar to `user@ubuntu~f321a` for your terminal session.
 4. START LEARNING!

 ## Labs

### Lab 1 - Configuring App Service Configurable
### Lab 2 - App Service Configurable - Ingest custom types via REST
### Lab 3 - Using the App Functions SDK w/ custom functions
### Lab 4 - App Functions SDK - Commanding a Device Service
