# 2.0 Introduction (Structs)

Current problem with the cards project: unable to directly access the suit/value of a specific card, as they are all stored in a single string of information.

## Struct

* A data structure in Go which is a **collection of properties** that are related together.
    - In cards, we could have created a struct of type card, that has fields of `suit` and `value` that are both strings
    - You can create as many "instances" of the struct as you wish
    - Sample data
        * A person has
            - A first name of `string`
            - A last name of `string`
        * Define the contents of the struct outside of main
            ```go
            type person struct {
                firstName string
                lastName string
            }
            ```
        * Create an "instance" of the struct
            - In order of field definition (not recommended because it may cause trouble)
                ```go
                alex := person{"Alex", "Anderson"}
                ```
            - Specifc definition
                ```go
                john := person{firstName: "John", lastName: "Doe"}
                ```
            - Zero value assignment
                * A struct is created and a **zero value** of the specific data type will be assigned to the fields
                    - Zero values: default values of an uninitialized variable
                        * `string` - `""`
                        * `int` - `0`
                        * `float` - `0`
                        * `bool` - `false`
                ```go
                var joe person
                ```

## Manipulating Structs
* Printing structs
    - Directly printing
        ```go
        fmt.Println(john)
        // >> {John Doe}
        // printed out in order of fields
        ```
    - Printing out the field names and values of a struct
        ```go
        fmt.Printf("%+v", alex)
        // {firstName:Alex lastName:Anderson}
* Updating properties
    - The good 'ol dot structure
        ```go
        alex.firstName = "Alex"
        ```