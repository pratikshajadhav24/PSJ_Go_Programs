package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

// Define a structure matching the XML format
type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
    City    string   `xml:"city"`
}

func main() {
    // Open the XML file
    file, err := os.Open("data.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Read the file content
    byteValue, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Unmarshal the XML data into the struct
    var person Person
    err = xml.Unmarshal(byteValue, &person)
    if err != nil {
        fmt.Println("Error unmarshalling XML:", err)
        return
    }

    // Display the structure
    fmt.Println("Person Details:")
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    fmt.Println("City:", person.City)
}
