package main
import (
	"context"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/mmildd_s/app/controllers"
	"github.com/mmildd_s/app/ent"
)
type Patients struct {
	Patient []Patient
}

type Patient struct {
	PatientName  string
	PatientAge  string
	PatientWeight  string
	PatientHeight  string
	PatientPrefix  string
	PatientGender  string
	PatientBlood  string
}

type SchemeTypes struct {
	SchemeType []SchemeType
}

type SchemeType struct {
	SchemeTypeName string
}

type Funds struct {
	Fund []Fund
}

type Fund struct {
	FundName string
}

type Certificates struct {
	Certificate []Certificate
}

type Certificate struct {
	CertificateName string
}

type Medicals struct {
	Medical []Medical
}

type Medical struct {
	MedicalName  string
	MedicalEmail  string
	MedicalPassword  string
	MedicalTel  string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	client, err := ent.Open("sqlite3", "file:person.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	v1 := router.Group("/api/v1")
	controllers.NewPatientController(v1, client)
	controllers.NewSchemeTypeController(v1, client)
	controllers.NewFundController(v1, client)
	controllers.NewCertificateController(v1, client)
	controllers.NewMedicalController(v1, client)
	controllers.NewCoveredPersonController(v1, client)
	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"สุมารี นามใจ","30","55","157","นาง","หญิง","B" },
			Patient{"สมชาติ มีใจ","28","67","178","นาย","ชาย","O" },
			Patient{"นานมี อารมณ์ดี","22","50","162","นางสาว","หญิง","A" },
		},
	}
	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetPatientName(p.PatientName).
			SetPatientAge(p.PatientAge).
			SetPatientWeight(p.PatientWeight).
			SetPatientHeight(p.PatientHeight).
			SetPatientPrefix(p.PatientPrefix).
			SetPatientGender(p.PatientGender).
			SetPatientBlood(p.PatientBlood).
			Save(context.Background())
	}
	// Set SchemeType Data
	schemeTypes := SchemeTypes{
		SchemeType: []SchemeType{
			SchemeType{"สิทธิการรักษาพยาบาลของข้าราชกาล"},
			SchemeType{"สิทธิประกันสังคม"},
			SchemeType{"สิทธิหลักประกันสุขภาพ"},
		},
	}
	for _, st := range schemeTypes.SchemeType {
		client.SchemeType.
			Create().
			SetSchemeTypeName(st.SchemeTypeName).
			Save(context.Background())
	}
	// Set Fund Data
	funds := Funds{
		Fund: []Fund{
			Fund{"ด้านการรักษาพยาบาล"},
			Fund{"ด้านการส่งเสริมสุขภาพ และป้องกันโรค"},
			Fund{"ด้านการฟื้นฟูสุขภาพ / แพทย์แผนไทย"},
		},
	}
	for _, f := range funds.Fund {
		client.Fund.
			Create().
			SetFundName(f.FundName).
			Save(context.Background())
	}

	// Set Certificate Data
	certificates := Certificates{
		Certificate: []Certificate{
			Certificate{"รับใบรับรองแพทย์"},
			Certificate{"ไม่รับใบรับรองแพทย์บ"},
		},
	}
	for _, ce := range certificates.Certificate {
		client.Certificate.
			Create().
			SetCertificateName(ce.CertificateName).
			Save(context.Background())
	}
	// Set Medicals Data
	medicals := Medicals{
		Medical: []Medical{
			Medical{"สมศรี ภาคภูมิ","somsee@hotmai.com","123","0989987653" },
			Medical{"มีใจ มะมี","mejai@hotmai.com","123","0897563456" },
			Medical{"ถังแก๊ส เดินได้","tuggas@hotmai.com","123","0893456324" },
		},
	}
	for _, mm := range medicals.Medical {
		client.Medical.
			Create().
			SetMedicalName(mm.MedicalName).
			SetMedicalEmail(mm.MedicalEmail).
			SetMedicalPassword(mm.MedicalPassword).
			SetMedicalTel(mm.MedicalTel).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}