// package alexa exports gopherjs bindings to the node library "alexa-sdk", as
// well as utility functions. This lets you write code in Go, and then turn it
// into a javascript module suitable for running on AWS Lambda.
package alexa

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

var handlers = map[string]interface{}{}

// These are the standard Amazon Intents retrieved from
// https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents
const (
	HELP_INTENT        = "AMAZON.HelpIntent"
	CANCEL_INTENT      = "AMAZON.CancelIntent"
	LOOP_OFF_INTENT    = "AMAZON.LoopOffIntent"
	LOOP_ON_INTENT     = "AMAZON.LoopOnIntent"
	NEXT_INTENT        = "AMAZON.NextIntent"
	NO_INTENT          = "AMAZON.NoIntent"
	PAUSE_INTENT       = "AMAZON.PauseIntent"
	PREVIOUS_INTENT    = "AMAZON.PreviousIntent"
	REPEAT_INTENT      = "AMAZON.RepeatIntent"
	RESUME_INTENT      = "AMAZON.ResumeIntent"
	SHUFFLE_OFF_INTENT = "AMAZON.ShuffleOffIntent"
	SHUFFLE_ON_INTENT  = "AMAZON.ShuffleOnIntent"
	RESTART_INTENT     = "AMAZON.StartOverIntent"
	STOP_INTENT        = "AMAZON.StopIntent"
	YES_INTENT         = "AMAZON.YesIntent"
)

// HandleFunc registers a function as the handler for a given request.
func HandleFunc(name string, handlerFunc func(Response)) {
	handlers[name] = toFunction(handlerFunc)
}

// ListenAndServe blocks and serves the requests that have been set up by
// calling HandleFunc.
func ListenAndServe(module *js.Object) {
	module.Get("exports").Set(
		"handler",
		func(event, context, callback *js.Object) {
			_alexa := js.Global.Call("require", "alexa-sdk")

			var alexaResponse = _alexa.Call(
				"handler", event, context)
			alexaResponse.Call("registerHandlers", handlers)
			alexaResponse.Call("execute")
		})
}

// A Response is used by the Handler Function to construct a ...response?
type Response struct {
	inner *js.Object
}

// Redirect handles the current method with the redirected method instead.
func (h Response) Redirect(method string) {
	h.emit(method)
}

// Tell emits a single message to the client, and then closes the session.
func (h Response) Tell(message string) {
	h.emit(":tell", strings.Replace(strings.Replace(message, "\n", " ", -1), "\t", "", -1))
}

// Ask emits a prompt to the client, and keeps the session open. If the prompt
// is not responded to, the reprompt will be issued.
func (h Response) Ask(prompt, reprompt string) {
	h.emit(":ask", prompt, reprompt)
}

func (h Response) emit(args ...interface{}) {
	h.inner.Call("emit", args...)
}

func (h Response) On() {}

func (h Response) EmitWithState() {}

func (h Response) State() {}

func (h Response) I18n() {}

func (h Response) Locale() {}

func (h Response) Localize() {}

func (h Response) Event() {}

func toFunction(fn func(Response)) *js.Object {
	return js.MakeFunc(
		func(this *js.Object, arguments []*js.Object) interface{} {
			fn(Response{inner: this})
			return nil
		})
}

// Break returns a series of one or more <break> ssml tags to get the required length of pause.
func Break(t time.Duration) string {
	buf := bytes.Buffer{}
	for ; t > 10*time.Second; t -= 10 * time.Second {
		buf.WriteString(fmt.Sprintf(`<break time="%s"/>`, 10*time.Second))
	}
	buf.WriteString(fmt.Sprintf(`<break time="%s"/>`, t))
	return buf.String()
}

// TODO(danver): Fill these out eventually.
// attributes: this._event.session.attributes,
// context: this._context,
// name: eventName,
// isOverridden:  IsOverridden.bind(this, eventName),
// response: ResponseBuilder(this)
