# Tutorial 3

## Mouse Move

> [On Github: Tutorial 3](tutorial-03-MouseMove)

For clarity, some of the code will be hidden.  View the full source code at the repository.

We will skip over material that has been covered in previous tutorials.

&nbsp;

We can get nodes to listen to mouse move events, by using `bl.OnMouseMove(...)`

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

			bl.OnMouseMove(func(e *bl.MouseMoveEvent) {
				fmt.Println("Mouse moved on ", e.Target.Id, "at", e.X, e.Y)
			})

			border.Fill(0, 50, 0)
			border.Wire()
		}
		bl.End()
	}
	bl.End()
}
```

Lets focus on

```
	bl.OnMouseMove(func(e *bl.MouseMoveEvent) {
		fmt.Println("Mouse moved on ", e.Target.Id, "at", e.X, e.Y)
	})
```

We can determine exactly what was done by using the information from the event.

And we are done!

