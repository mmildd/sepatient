package schema
 
import (
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent"
)
// CoveredPerson holds the schema definition for the CoveredPerson entity.
type CoveredPerson struct {
    ent.Schema
}
// Fields of the CoveredPerson.
 func (CoveredPerson) Fields() []ent.Field {
    return []ent.Field{
  
    }
 }

 //Edges of the CoveredPerson.
 func (CoveredPerson) Edges() []ent.Edge {
   return []ent.Edge{
      edge. From("Patient",Patient.Type).Ref("Patient_CoveredPerson").Unique(),
      edge. From("SchemeType",SchemeType.Type).Ref("SchemeType_CoveredPerson").Unique(),
      edge. From("Fund",Fund.Type).Ref("Fund_CoveredPerson").Unique(),
      edge. From("Certificate",Certificate.Type).Ref("Certificate_CoveredPerson").Unique(),
   }
 }