# raytrace

A raytracing program written in Go.

Created following the book [Ray Tracing in One Weekend](http://in1weekend.blogspot.com/2016/01/ray-tracing-in-one-weekend.html) by Peter Shirley (original source code written in C++).

![Output of the final program](/docs/final.png)

## Getting Started

Build the binary, run the program to generate an image, and open it in an image viewer.

```bash
go build -o ./bin/raytracer ./src/
./bin/raytracer > output.ppm
open output.ppm
```

## Resources

* [Kindle version of book](https://read.amazon.com/?asin=B01B5AODD8)
