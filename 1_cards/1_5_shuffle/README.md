# 1.5 Shuffle
12/25/2018

**Goal: Shuffle the order of elements in a deck, which is a slice of strings**

## Pseudocode

* For each index, card in cards
    - Generate a random number between 0 and len(cards) - 1
    - Swap the current card and the card at cards(randNum)

## Implementation

* `len(a)` is used to return the size (length) of the slice `a`.
* To swap the values of two variables or slice elements, use the swapping syntax
    ```go
    a, b = b, a
    ```
* Random Number Generation
    - Using the `Intn` function from the `math/rand` standard package to generate random numbers.
        ```go
        func Intn(n int) int
        ```
        > Intn returns, as an int, a non-negative **pseudo-random** number in [0,n) from the default Source. It panics if n <= 0.

        * It's using the same seed every time it's called, therefore generating the same sequence every time. **Not really random.**
    - Type `Rand` is a source of random numbers - reinventing the wheel, but necessary for "real" randomness
        ```go
        // use current time as the seed of random generator
        // time.Now().UnixNano() retuns an int64 type of current time
        source := rand.NewSource(time.Now().UnixNano())

        // create a new random generator with the seed
        r := rand.New(source)

        // use the generator with the new seed instead of the pseudo generator
        r.IntN(n)
        ```