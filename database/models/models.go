// nx/database/models will provide a way to define data models programmatically.
// Unlike the well-defined datatypes you create for packages with a specific purpose,
// models will allow you to define a data model and will use the text/template package
// in order to generate the model definitions before compile time.
// In order for this to work...a conversion program will have to be build in the cmd/
// folder which must be built and run before the primary program is built as it will rely
// on the generated data models.
package models
