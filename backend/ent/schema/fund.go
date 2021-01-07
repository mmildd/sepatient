package schema
 
import (
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent"
)
// Fund holds the schema definition for the Fund entity.
type Fund struct {
    ent.Schema
}
// Fields of the Fund.
 func (Fund) Fields() []ent.Field {
    return []ent.Field{
        field.String("Fund_Name").NotEmpty(),
    }
 }
 //Edges of the Fund.
  func (Fund) Edges() []ent.Edge {
   return []ent.Edge{
		edge. To("Fund_CoveredPerson",CoveredPerson.Type).StorageKey(edge.Column("Fund_id")),

 }
}
