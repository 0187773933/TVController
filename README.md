# TV Controller

- Supported :
	- LG
	- Vizio
	- HDMICEC
	- IR :
		- https://irdroid.com/irdroid-usb-ir-transceiver
		- https://irdroid.com/irdroid-wifi-version-3-0
		- https://www.lirc.org/software.html
		- https://github.com/Irdroid/Irdroid-USB/blob/master/app/src/main/java/com/microcontrollerbg/usbirtoy/IrToy.java
		- https://github.com/ww24/lirc-web-api/blob/master/lirc/client.go
		- https://github.com/ww24/lirc-web-api/blob/master/lirc/test/lircd.conf
		- https://lirc-remotes.sourceforge.net/remotes-table.html
		- https://github.com/Irdroid/USB_IR_Toy
		- https://irdroid.com/2016/10/how-to-turn-your-raspberry-pi-into-a-fully-functional-infrared-remote-control/
		- https://anibit.com/product/ptt08001
		- https://www.lirc.org/html/irtoy.html
		- https://github.com/irdroid/usb_infrared_transceiver
		- https://community.home-assistant.io/t/how-to-add-ir-tranmitter-to-enhance-your-audio-video-experince/446033

- Todo :
	- LG Apps
		- `npm install -g @webosose/ares-cli`
		- https://webostv.developer.lge.com/develop/tools/cli-installation
		- https://www.webosose.org/docs/reference/ls2-api/com-webos-service-applicationmanager/
		- https://github.com/home-assistant/core/blob/dev/homeassistant/components/webostv/const.py
		- https://webostv.developer.lge.com/develop/getting-started/build-your-first-web-app
		- https://webostv.developer.lge.com/develop/tools/sdk-introduction
		- https://webostv.developer.lge.com/develop/tools/cli-introduction
		- https://webostv.developer.lge.com/develop/tools/simulator-introduction
		- https://webostv.developer.lge.com/develop/guides/js-service-usage
		- https://github.com/webOS-TV-app-samples/webOSTVJSLibrary
		- https://webostv.developer.lge.com/develop/references/webostvjs-webos#request
	- Samsung
		- https://github.com/McKael/samtv
		- https://github.com/mgoff/go-samsung-exlink
		- https://github.com/rainu/samsung-remote-mqtt
		- https://github.com/rainu/samsung-remote
	- Sony
	- Hisense
		- https://github.com/Krazy998/mqtt-hisensetv
		- https://github.com/yosssi/gmq
	- TCL
	- Westinghouse
	- https://a.co/d/1CtCSsp

## Misc

- both the firecube and xbox seem to use HDMI-CEC for power control, and then IR for everything else