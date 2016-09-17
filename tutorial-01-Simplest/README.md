# Tutorial 1

## Simplest Bellina application

This is the simplest possible Bellina application.

> [On Github: Tutorial 1](https://github.com/amortaza/go-bellina-tutorials/tree/master/tutorial-01-Simplest)

[[tutorial-01-Simplest/tutorial-01-a.jpg]]

For clarity, some of the code will be hidden.  View the full source code at the repository.

```
func init() {
	runtime.LockOSThread()
}
```

Our GLFW library requires the above.

&nbsp;

```
func main() {
	bl.Start( hal_g5.New(), 1280, 1024, "Bellina v0.9", initialize, tick, uninitialize )

	fmt.Println("bye!")
}
```

Bellina is started with call to `bl.Start(...)`

**1st parameter** is an implementation of the `bl.HAL` interface which is what feeds Bellina the keyboard / mouse events and also provides a graphics object which allows rendering.

Here we use a reference implementation called `amortaza/go-hal-g5`

**2nd and 3rd parameters** we set the application window to dimension 1280 x 1024.

**4th parameter** the title reads "Bellina v0.9".

**5th parameter** is a callback function called `initialize`.

This function is called back by Bellina after the main application window has been created / GLFW and OpenGL have been initialized.

In this code you can put any application specific initialization that are needed.

**6th parameter** is tick callback function (referred to as the `user-tick` function).

This function gets called at every loop for your user application to update state.  Its the main loop of your application.

**7th parameter** is uninitialize callback function.  Once the application window is closed, the uninitialize callback function is called. It is after this function returns that Bellina cleans up (e.g. unloading and freeing up various libraries and resources).

> If you don't need an initialize and/or uninitialize callback functions, you may pass `nil`.

&nbsp;

```
func initialize() {
	fmt.Println("Initializing!")

	go_dark_ux.Init()
}

func uninitialize() {
	fmt.Println("Uninitializing!")

    go_dark_ux.Uninit()
}
```

Since Bellina does not do any rendering, the application would look blank.  So we are going to make use of a `border` plugin that draws borders and can fill a node with various colors.  The details are left for later, but for now we need to import and initialize the `go_dark_ux` library that is in `github.com/amortaza/go-dark-ux`.

Bellina does not impose any interface rules on plugins and libraries.  Each plugin will have its own needs for setup and teardown, so make sure you follow each plugin's.

> Bellina convention is that a plugin, if it needs one, to expose `Init()` and `Uninit()` methods with whatever parameters are relevant.

&nbsp;

```
func tick() {

	bl.Root()
	{
		bl.Dim( bl.Window_Width, bl.Window_Height)

		border.Fill(50, 0, 0)
		border.Wire()
	}
	bl.End()
}
```

Finally we get to the main loop of our application, the `user-tick` callback function.

Each call to this function is called a `frame`.  Just like, say, a video game frame.

> Bellina node tree is destroyed every frame.  The Bellina node tree is re-created in the tick function.  With a call to `bl.Root()`

Here we size the `root` node as the same dimension as the application window.

There is no need to position the root node because it is *always* at (0, 0).

We use the border plugin to Fill the node with color and to draw a border around it.

Do not concern yourself with how plugins or how border works at this time.

&nbsp;

Lets add a child node.

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
			bl.Dim(90, 60)

			border.Fill(0, 50, 0)
			border.Wire()
		}
		bl.End()
	}
	bl.End()
}
```

Now that we are experts, lets add more child nodes and even a grandchild.

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
			bl.Dim(90, 60)

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

			bl.Div()
			{
				bl.Id("grand child")
				bl.Pos(20, 40)
				bl.Dim(120, 90)

				border.Fill(30, 80, 10)
				border.Wire()
			}
			bl.End()
		}
		bl.End()
	}
	bl.End()
}
```

We are noticing that this is becoming a long function.
Since this is just regular Go, there is no reason why this cannot be broken up into functions to manage code better.

Try putting the child creation pieces in separate functions (the answer is below, but is not in the Github repo).

We will see later that breaking out Bellina code into functions is the basis of plugins.

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
			bl.Dim(90, 60)

			border.Fill(0, 50, 0)
			border.Wire()
		}
		bl.End()

		sibling()
	}
	bl.End()
}

func sibling() {
    bl.Div()
	{
		bl.Id("child 2")
		bl.Pos(40, 80)
		bl.Dim(290, 260)

		border.Fill(0, 10, 80)
		border.Wire()

        grandChild()
	}
	bl.End()
}

func grandChild() {
	bl.Div()
	{
		bl.Id("grand child")
		bl.Pos(20, 40)
		bl.Dim(120, 90)

		border.Fill(30, 80, 10)
		border.Wire()
	}
	bl.End()
}
```

And we are done!

