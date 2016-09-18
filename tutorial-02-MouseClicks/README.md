# Tutorial 2

## Mouse Clicks

> [On Github: Tutorial 2](https://github.com/amortaza/go-bellina-tutorials/tree/master/tutorial-02-MouseClicks)

For clarity, some of the code will be hidden.  View the full source code at the repository.

We will skip over material that has been covered in previous tutorials.

&nbsp;

We can get nodes to listen to mouse click events, by using `bl.OnMouseButton(...)`

```
func tick() {

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

			bl.OnMouseButton(func(e *bl.MouseButtonEvent) {
				fmt.Println("Mouse clicked on ", e.Target.Id, "at", e.X, e.Y)
			})

			border.Fill(0, 50, 0)
			border.Wire()
		}
		bl.End()

		bl.Div()
		{
			bl.Id("child 2")
			bl.Pos(40, 80)
			bl.Dim(290, 260)

			border.Fill(0, 10, 80)
			border.Wire()

			bl.OnMouseButton(func(e *bl.MouseButtonEvent) {
				fmt.Println("Clicking on ", e.Target.Id, "at", e.X, e.Y)
			})

			bl.Div()
			{
				bl.Id("grand child")
				bl.Pos(20, 40)
				bl.Dim(120, 90)

				border.Fill(80, 10, 80)
				border.Wire()
			}
			bl.End()
		}
		bl.End()
	}
	bl.End()
}
```

Note that *only* `child` and `child 2` are listening to mouse clicks and not `grand child`.

Lets focus on

```
	bl.OnMouseButton(func(e *bl.MouseButtonEvent) {
		fmt.Println("Clicking on ", e.Target.Id, "at", e.X, e.Y)
	})
```

The callback function is called and prints information about what was clicked.

Notice that a click generates *two* calls to the callback.  

> That is because the callback is called both on the *down* press and the *up* release.

We can determine exactly what was done by using the information from the event.

Say we only want to print on the *up* release.

```
	bl.OnMouseButton(func(e *bl.MouseButtonEvent) {
		if e.ButtonAction == bl.Button_Action_Up {
			fmt.Println("Mouse button released", e.Target.Id, "at", e.X, e.Y, "type", e.ButtonAction)
		}

		if e.ButtonAction == bl.Button_Action_Down {
			fmt.Println("Mouse button pressed down", e.Target.Id, "at", e.X, e.Y, "type", e.ButtonAction)
		}
	})
```

Now say we want to determine which button was pressed.

```
	bl.OnMouseButton(func(e *bl.MouseButtonEvent) {
		if e.Button == bl.Mouse_Button_Left {
			fmt.Println("Left Mouse button", e.Target.Id, "at", e.X, e.Y, "type", e.ButtonAction)
		}
	})
```

And we are done!

