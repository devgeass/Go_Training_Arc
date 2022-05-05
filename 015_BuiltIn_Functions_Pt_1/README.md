# #15 BuiltIn Functions Part 1

## Math

Math functions allow performing mathematical operations including logarithmic and trignometric functions. You can access these functions using `math.func(args)`.  Here's a list:

| Function  | Syntax         | Returns                         |
| --------- | -------------- | ------------------------------- |
| Abs       | Abs(x)         | Absolute value of x             |
| Ceil      | Ceil(x)        | Least Int value >= x            |
| Floor     | Floor(x)       | Greater Int value <=x           |
| Log       | Log(x)         | Natural log of x                |
| Log10     | Log10(x)       | Log base 10 of x                |
| Pow       | Pow(x,y)       | x**y                            |
| Pow10     | Pow10(x)       | 10**x                           |
| Sqrt      | Sqrt(x)        | Square root of x                |
| Max       | Max(x,y)       | Maximum between x and y         |
| Min       | Min(x,y)       | Min between x and y             |
| Sin       | Sin(x)         | Sine of radian x                |
| Cos       | Cos(x)         | Cosine of radian x              |
| Tan       | Tan(x)         | Tangent of radian x             |
| Mod       | Mod(x,y)       | Floating-point remainder of x/y |
| Remainder | Remainder(x,y) | IEEE 754 remainder of x/y       |
| Exp       | Exp(x)         | e**x                            |
| Inf       | Inf(x)         | +INF for x>=0; -INF for x<0     |

Math package also provides some standard constants.

| Constant | Syntax   | Value             |
| -------- | -------- | ----------------- |
| PI       | math.Pi  | 3.141592653589793 |
| PHI      | math.Phi | 1.618033988749895 |
| e        | math.E   | 2.718281828459045 |

## Random

Random integers or floats can be produced using `math/rand` package. You can access these functions using `math.rand(args)`

| Function | Syntax    | Returns                                                            |
| -------- | --------- | ------------------------------------------------------------------ |
| Int      | Int()     | a non-negative pseudo-random int                                   |
| Intn     | Intn(x)   | a pseudo-random int from interval [0,x)                            |
| Float32  | Float32() | a pseudo-random float32 from interval [0.0,1.0)                    |
| Float64  | Float64() | a pseudo-random float64 from interval [0.0,1.0)                    |
| Perm     | Perm(x)   | a slice of x ints containing pseudo-random int from interval [0,n) |

That's a wrap on day 15.