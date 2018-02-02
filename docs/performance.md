# Performance

Once finished with the book, I set out to improve the performance of the program. All times below are from a 2015 MacBook Pro.

Very few pointers (200 x 100 x 20 samples) [commit 3765884]

```bash
time ./bin/raytracer > output.ppm

# real	0m34.895s
# user	0m36.306s
# sys	0m0.940s
```

After pointers (200 x 100 x 20 samples) [commit 3ef4c49]

```bash
time ./bin/raytracer > output.ppm

# real	0m23.360s
# user	0m24.260s
# sys	0m0.691s
```

After pointers (400 x 200 x 200 samples) [commit 3ef4c49]

```bash
time ./bin/raytracer > output.ppm

# real	16m34.618s
# user	17m26.047s
# sys	0m25.205s
```
