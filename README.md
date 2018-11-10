# HueScenes4openHAB

This is a little tool for generating openHAB config files for dealing with Hue Scenes, which
is not supported by the Hue binding. This way of dealing with scenes works nice for me. However,
I'm open for suggestions on how to improve this, feel free to open issues or pull requests.

## What do you get?

You get config files for Items, Rules and Maps. The idea is this:

- There will be a String Item for every group / room, which represents the scene to activate.
- There will be a Rule for every group / room, which reacts on the relevant item state change,
and uses the Hue HTTP API for actually activate the scene.
- Finally there will be a Map, which translates between human readable scene names and their Hue IDs.

TODO add examples

## Usage

In this first version you need to prepare access to your Hue bridge a bit:

- Find the IP address of the Hue bridge and write it down (check your router, or find it in the Hue app)
- Create a Hue user:
	 - Visit `http://<bridge-ip>/debug/clip.html` in your browser
	 - enter `/api/` in the "URL" field
	 - enter `{"devicetype":"openhab#huescenes"}` in the "Message Body" field (feel free to use another value,
	 but it has to be in format "abc#def")
	 - Push the physical button your Hue bridge
	 - Click the "POST" button within 30 seconds
	 - Write down the username (the long char/number combination, without quotes)
	 
Now download the latest release of HueScenes4openHAB here: TODO

And finally generate the config files with:

	HueScenes4openHAB --hueIP <bridge-ip> --hueUser <username>
	
The filenames wil all start with `huescenes`, you can override this with the `--filename <yourNewFilename>` paramater.

If something goes wrong, enable debugging with the `--debug` flag, in order to get more information.

Now copy the files to your openHAB installation, and configure your UI with the new Items. You need to use
`Selection`s, and the choices need to be the scene names of the corresponding group / room. You can add another choice
named "Off", which will switch off all light in that group / room.

## Roadmap

These items are on my todo list, but I have no estimation on when (and tbh, if at all) I will get to this:

- Detect the hue bridge automatically
- Create a user automatically (only the physical button press on the bridge needed)
- Optionally create Items for the single lights

## License

Copyright 2018 Marc Sluiter

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.