## Validator Format Message  



## example 
```go

v := validator.InitlizeValidator()
jsonData := []byte(`{
    "name": "alaa aqeel",
    "username": "al",
    "password": "123",
    "password_confirmation": "123456"
}`)

var user CreateUserDto
if err := json.Unmarshal(jsonData, &user); err != nil {
    fmt.Println("Invalid JSON:", err)
    return
}

errors := v.Struct(user)
if errors.HasErrors() {
    fmt.Println(errors)
    // Output:
    // {
    //     "name": ["required", "min", "max"],
    //     "username": ["required", "min", "max"],
    //     "password": ["required", "min", "max"],
    //     "password_confirmation": ["eqfield"]
    // }
}

```