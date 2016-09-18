# Tutorial 6

## Click Plugin

> [On Github: Tutorial 6](tutorial-06-Plugin-Click)

For clarity, some of the code will be hidden.  View the full source code at the repository.

We will skip over material that has been covered in previous tutorials.

&nbsp;

We have been introduced to many of the concepts needed to be able to create our first plugin - `Click`.

As a plugin designer, there is no preconditions on how things should be laid out, except to make usage of plugin as useful and easy as possible.

First, lets design what the plugin does, and how it is used.

## Click plugin API design

* `Click` maybe used without instantiating any object

* `Click` will call a callback function when a node is *clicked*

* *clicked* is defined as a mouse-down event followed by a mouse-up event on the same node

* Pressing mouse button on a node, but releasing the mouse button on a different node is *not* a click

* To make the plugin more useful, we can provide `life-cycle` methods
	* When the click life-cycle starts, there will be a callback
    * When the click action completes successfully, there will be a callback
    * When the click action is incomplete (ie. see *clicked* definition above), there will be a callback

* Because the simple case of a successful click notification is the 90% case, we will make it easier to use

* The full life-cycle setup, is more complicated and far less common so we will make that a separate function

Here is what the API will look like:

```
click.On( callback func(interface{}) )

click.On_WithLifeCycle( successCallback, onButtonDown, onButtonUpAndMiss func(interface{})
```

And we are done!
