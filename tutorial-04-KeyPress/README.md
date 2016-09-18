# Tutorial 4

## Key Press

> [On Github: Tutorial 4](https://github.com/amortaza/go-bellina-tutorials/tree/master/tutorial-04-KeyPress)

For clarity, some of the code will be hidden.  View the full source code at the repository.

We will skip over material that has been covered in previous tutorials.

&nbsp;

This is going to be a bit different than previous tutorials since, unlike the mouse events, Bellina does not have a ~~bl.OnKey~~ method.  This is because to have such a method, a multi-step process has to occur.
* The node has to have the keyboard focus 
* User selects node for focus through tabbing or clicking or some other method
* As keys are pressed, the even is directed to node with the focus

This is too high-level for Bellina.  
The above activities have to be coordinated by a *Keyboard Focus Plugin* which can be customized for various needs.  

Instead Bellina fires events that can be registered for by one or more callbacks.

Lets try that below.

Looking under `Events API` in the Bellina API reference we see 2 choices for registering for events.
* Short term
* Long term

The difference is that *Short Term* is only valid during this frame.  The next frame, the node and the registration will have disappeard - and will need to be re-created.

Long term is usually used for plugins so we will go with Short term.

The `bl.RegisterShortTerm(eventType string, cb func(Event))` needs an `eventType`.  

The event type of a key-event is defined in `bl.EventType_Key`.

Putting it all together.

```
bl.RegisterShortTerm(bl.EventType_Key, func(event bl.Event) {
})
```

`event` is of type `bl.Event` interface.  However, it really is a `bl.KeyEvent` object.

```
bl.RegisterShortTerm(bl.EventType_Key, func(event bl.Event) {
    keyEvent := event.(*bl.KeyEvent)
})
```

Also note that the registration of the event is not within a node context.  The registration is independent of any node.  It simply registers a callback to be called whenever a key is pressed.

```
func tick() {

	bl.RegisterShortTerm(bl.EventType_Key, func(event Event) {
		fmt.Println("Mouse clicked on ", e.Target.Id, "at", e.X, e.Y)
	})

	bl.Root()
	{
		bl.Dim( bl.Window_Width, bl.Window_Height)

		border.Fill(50, 0, 0)
		border.Wire()

		bl.Div()
		{
			bl.Id("child")
			bl.Pos(20, 40)
			bl.Dim(130, 160)

			border.Fill(0, 50, 0)
			border.Wire()
		}
		bl.End()
	}
	bl.End()
}
```

The registration generates *two* calls to the callback whenever a key is pressed and released.

> That is because the callback is called both on the *down* press and the *up* release.

We can determine exactly what was done by using the information from the event.

To examine the *up* release:

```
bl.RegisterShortTerm(bl.EventType_Key, func(event bl.Event) {
	keyEvent := event.(*bl.KeyEvent)

	if keyEvent.Action == bl.Button_Action_Down {
		fmt.Println("Down", keyEvent.Key)
	}

	if keyEvent.Action == bl.Button_Action_Up {
		fmt.Println("Down", keyEvent.Key)
	}
})
```

The `Alt`, `Ctrl`, and `Shift` properties of the event are `true` when these control buttons are pressed at the time of the event generation - `false` if the control buttons are not pressed.

And we are done!

