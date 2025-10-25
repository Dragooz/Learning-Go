Completed: https://go.dev/tour

1. := vs var

-   :=
    i. can only be used INSIDE function
    ii. Cannot be redeclare with :=, but can be reassigned

2. what is uint vs int

-   int = Signed Integer = can be + or -
-   uint = Unsigned Integer = can be + only

3. why pointer exist in Go? What's the usage? Pros and cons? When should I use it?

-   -   means when you wanna call the value out: E.g. fmt.Println(\*ptr) -> result in value; If you print ptr, it will result in the memory address, e.g. 142090...
    *   OR when you wanna declare a type, E.g. var ptr \*int >> Means ptr will store the address of a variable, with type int.
-   & means when you wanna get the address of the variable, E.g. x:=42; ptr := &x; ptr is now the address of x.

4. fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s), what is the fstring syntax here?

-   This is Go format verbs, not Python f-strings

5. what's the benefits of having %d and %v? can't %v serve all the purpose

-   For formatting control, E.g. %.1f
-   For bug discovery, E.g. age := "25", but fmt.Printf("%d", age) > Compiler warning or runtime error!

6. why uint8 instead of int8

-   uint8 means: "this variable can only store values from 0 to 255"
-   int means: "This variable can store values from -127 top 127"
-   If declared uint8(256) > It will hit error!
-   uint8 on existing var, will act as modulus

7. func (v _Vertex) Scale(f float64) {v.X = v.X _ f} vs func (v Vertex) Scale(f float64) {v.X = v.X \* f}

-   with \*, the underlying vertex variable will update the variables.
-   without \*, the underlying vertex variable will not update the variable.
