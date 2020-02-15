# CLI for Xiaomi Bulb 2

The app provides command-line interface to control MiBulb 2. Implementation is based on official network protocol: https://www.yeelight.com/en_US/developer

## Setup
- enable LAN control https://www.yeelight.com/faqs/lan_control
- navigate to the package root dir and `go build -o <utilName>`

Here `<utilName>` represents future util's exec name, can be anything a user prefer
- move the generated executable `<utilName>` to any dir and add that dir into PATH env var

Once done, the util is accessible from any terminal working dir
- now do `<utilName> discover` for initial setup
- finally, type `<utilName>` without args to browse available commands