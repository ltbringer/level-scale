package integrationtest

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"level-scale/dbmanager"
	"level-scale/settings"
	"net/http"
	"os"
)

var serviceHost string
var _ = BeforeSuite(func() {
	settings.Init()
	serviceHost = os.Getenv("SERVICE_HOST")
})

var _ = Describe("User Registration API", func() {
	var db *sql.DB
	var err error

	BeforeEach(func() {
		dsn := dbmanager.CreateDsn(settings.DbConfig)
		db, err = sql.Open("postgres", dsn)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		_, err = db.Exec(`TRUNCATE TABLE users RESTART IDENTITY CASCADE;`)
		Expect(err).NotTo(HaveOccurred())
		err = db.Close()
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}
	})

	It("should register a new user", func() {
		baseURL := fmt.Sprintf("http://%s:%d", serviceHost, settings.ServicePort)
		payload := map[string]interface{}{
			"email":    "tester@test.com",
			"password": "verysafepassword",
			"isSeller": false,
		}
		body, _ := json.Marshal(payload)
		res, err := http.Post(fmt.Sprintf("%s/%s", baseURL, "register"), "application/json", bytes.NewReader(body))
		Expect(err).NotTo(HaveOccurred())
		defer res.Body.Close()
		var respBody map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&respBody)
		Expect(res.StatusCode).To(Equal(http.StatusCreated))
		Expect(respBody).To(HaveKey("id"))
		idFloat, ok := respBody["id"].(float64)
		Expect(ok).To(BeTrue())
		Expect(uint64(idFloat)).To(Equal(uint64(1)))
	})

})
