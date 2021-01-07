package schema
 
import (
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)
// SchemeType holds the schema definition for the SchemeType entity.
type SchemeType struct {
    ent.Schema
}
// Fields of the SchemeType.
func (SchemeType) Fields() []ent.Field {
    return []ent.Field{
        field.String("SchemeType_Name").NotEmpty(),
    }
 }
 //Edges of the SchemeType.
 func (SchemeType) Edges() []ent.Edge {
    return []ent.Edge{
		edge. To("SchemeType_CoveredPerson",CoveredPerson.Type).StorageKey(edge.Column("SchemeType_id")),
 }
 }