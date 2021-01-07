package schema
 
import (
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)
// Certificate holds the schema definition for the Certificate entity.
type Certificate struct {
    ent.Schema
}
// Fields of the Certificate.
func (Certificate) Fields() []ent.Field {
    return []ent.Field{
        field.String("Certificate_Name").NotEmpty(),
    }
 }
 //Edges of the Certificate.
 func (Certificate) Edges() []ent.Edge {
    return []ent.Edge{
		edge. To("Certificate_CoveredPerson",CoveredPerson.Type).StorageKey(edge.Column("Certificate_id")),
 }
 }