# 2.1 Manupulation
12/27/2018

## Printing structs
* Directly printing
    ```go
    fmt.Println(john)
    // >> {John Doe}
    // printed out in order of fields
    ```
* Printing out the field names and values of a struct
    ```go
    fmt.Printf("%+v", alex)
    // {firstName:Alex lastName:Anderson}
    ```

## Updating properties
* The good 'ol dot structure
    ```go
    alex.firstName = "Alex"
    ```
## Embedding Structs
* A struct can be embeded in another struct as a field type
    - Create all structs outside
        ```go
        type contactInfo struct {
            email string
            zipCode int
        }

        type person struct {
            contact contactInfo
            ...
        }
        ```
    - When constructing, declare the embeded struct as well (need struct name)
        ```go
        john := person{
	        firstName: "John",
            ...
            contact: contactInfo{
                email:   "email@gmail.com",
                zipCode: 94000,
		    },
        }
        ```
        * **When declaring a multi-line struct, all lines must end with a comma, even if it is the last line in the declaration!**