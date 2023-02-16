# Prime Number Tester
HTTP API to check if the list of given numbers are primes. Takes slice of integers, returns slice of booleans.

---

## Example

POST /

`[1,2,3,4,5]`


Responce

`[false, true, true, false, true]`

---
## Features
- Supports two prime number checkers: *built-in* one from [math/big](https://pkg.go.dev/math/big) package and *custom-built*. Custom *(used by default)* is faster for small values and from big package is faster for large values. The prime number checker can be changed using enviromental variable(see below).

---
## Quick start and settings
1. Clone git, compile code and run. (*Or run `main.go`)
2. Send POST request to \<your address\>:8000 *(by default used port 8000)*. The body should continue slice of integers.
3. Recieve slice of booleans.

*To change prime number checker - use ENV `CUSTOM_PRIME_CHECKER` (by default value is `true`, if you want to use one from math/big, set it to `false`)

---
## Authors

Bohdan Mykytenko