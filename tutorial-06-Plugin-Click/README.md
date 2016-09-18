# Tutorial 5

## Window Resize

> [On Github: Tutorial 5](tutorial-05-WindowResize)

For clarity, some of the code will be hidden.  View the full source code at the repository.

We will skip over material that has been covered in previous tutorials.

&nbsp;

Just like the `Key` events from the previous tutorials, Bellina fires events that can be registered for by one or more callbacks.

[See Tutorial 4 on the explanation of the Short and Long term listeners](tutorrial-04-KeyPress)

Lets try that below.

The `bl.RegisterShortTerm(eventType string, cb func(Event))` needs an `eventType`.  

The event type of a window-resize is defined in `bl.EventType_Window_Resize`.

Putting it all together.

```
bl.RegisterShortTerm(bl.EventType_Window_Resize, func(event bl.Event) {
})
```

`event` is of type `bl.Event` interface.  However, it really is a `bl.WindowResizeEvent` object.

```
bl.RegisterShortTerm(bl.EventType_Window_Resize, func(event bl.Event) {
    resizeEvent := event.(*bl.WindowResizeEvent)
})
```

Also note that the registration of the event is not within a node context.  The registration is independent of any node.  It simply registers a callback to be called whenever the application window is resized.

```
func tick() {

	bl.RegisterShortTerm(bl.EventType_Window_Resize, func(event bl.Event) {
		resizeEvent := event.(*bl.WindowResizeEvent)

		fmt.Println("Resized", resizeEvent.Width, resizeEvent.Height)
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

And we are done!
