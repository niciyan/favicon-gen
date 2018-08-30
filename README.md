favicon-gen is a simple favicon generator written in Go.

### Feature
favicon-gen creates simple shape favicon.
It aims for marking development web application.

* png format
* available shapes are circle, rectangle
* default size is 32x32

Examples are below:
* [rect favicon](./example/rect.png)
* [circle favicon](./example/circle.png)

### Usage
	$ favicon-gen > favicon.ico

### Option
	$ favicon-gen -h
	Usage of favicon-gen:
	  -circle
			paint circle for favicon
	  -l int
			length for favicon (default 32)
	  -t	use transparent background
	  -whiteBG
			white background

### TODO
many.
