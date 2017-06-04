# Gopherjs bindings to the Javascript Alexa SDK

It's easy to write a javascript Lambda function to power an Alexa skill using
the node library
[alexa-sdk](https://github.com/alexa/alexa-skills-kit-sdk-for-nodejs).

This library contains gopherjs bindings to alexa-sdk, so a developer can write a
lambda function in Go, and then "compile" it with gopherjs, link it against
alexa-sdk and have it run on AWS Lambda.

This library is only a temporary stopgap measure on the way to building
full-featured Alexa Skills in Go that run on AWS Lambda. Once I've properly
researched a way to have native Go running on Lambda, and have written a pure Go
implementation of the methods of this library, it should be very easy to port my
existing Alexa Skills written against this library to use the native engine.

Until then, this library already powers at least one Alexa skill published on
the Alexa Skills Store, so it works to some extent at least.

## Stability
None guaranteed.

## Contributing
No contributions are accepted at the moment. Please let me know if you'd like to
contribute, and I will inform you when this is in a better state.
