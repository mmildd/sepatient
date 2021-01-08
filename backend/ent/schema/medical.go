package schema
 
import (
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)

// User holds the schema definition for the Medical entity.
type Medical struct {
    ent.Schema
}

// Fields of the Medical.
func (Medical) Fields() []ent.Field {
    return []ent.Field{
		field.String("Medical_Name").NotEmpty(),
        field.String("Medical_Email").NotEmpty(),
        field.String("Medical_Password").NotEmpty(),
        field.String("Medical_Tel").NotEmpty(),
    }
 }

 
 // Edges of the Medical.
 func (Medical) Edges() []ent.Edge {
    return []ent.Edge{
       
    }
 }

 
