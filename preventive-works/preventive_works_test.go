package preventive_works_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/url"
)

var _ = Describe("PreventiveWorks", func() {
	var idPreventiveWork string

	Describe("Проверка preventive_works/", func() {
		Context("Проверка возвращаемых кодов", func() {
			//It("Возвращает 404", func() {
			//	resp, err := http.Get("http://localhost:8101/preventive_works/")
			//	Expect(err).To(BeNil())
			//	defer resp.Body.Close()
			//	Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			//})

			It("Возвращает 200", func() {
				resp, err := http.PostForm("http://localhost:8101/preventive_works/new_work", url.Values{
					"name_service": {"test"},
					"create_at":    {"2022-01-02 15:04:05"},
					"deadline":     {"2022-01-04 15:04:05"},
					"title":        {"test"},
					"description":  {"test"}})
				_, _ = http.Get("http://localhost:8101/preventive_works/")
				Expect(err).To(BeNil())
				defer resp.Body.Close()
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})

		})
	})

	Describe("Проверка preventive_works/new_work", func() {
		Context("Проверка возвращаемых кодов", func() {

			It("Возвращает 200", func() {
				resp, err := http.PostForm("http://localhost:8101/preventive_works/new_work", url.Values{
					"name_service": {"test"},
					"create_at":    {"2022-01-02 15:04:05"},
					"deadline":     {"2022-01-04 15:04:05"},
					"title":        {"test"},
					"description":  {"test"}})
				Expect(err).To(BeNil())
				defer resp.Body.Close()
				_, _ = resp.Body.Read([]byte(idPreventiveWork))

				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})

			It("Возвращает 500, если дата окончания больше даты создания", func() {
				resp, err := http.PostForm("http://localhost:8101/preventive_works/new_work", url.Values{
					"name_service": {"test"},
					"create_at":    {"2022-01-04 15:04:05"},
					"deadline":     {"2022-01-02 15:04:05"},
					"title":        {"test"},
					"description":  {"test"}})
				Expect(err).To(BeNil())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			})

			It("Возвращает 500, если дата введена неверно", func() {
				resp, err := http.PostForm("http://localhost:8101/preventive_works/new_work", url.Values{
					"name_service": {"test"},
					"create_at":    {"рп"},
					"deadline":     {"2022-01-02 15:04:05"},
					"title":        {"test"},
					"description":  {"test"}})
				Expect(err).To(BeNil())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			})
		})

		Describe("Проверка preventive_works/{id}", func() {

			Context("Проверка возвращаемых кодов", func() {

				It("Возвращает 200", func() {
					resp, err := http.Get("http://localhost:8101/preventive_works/" + idPreventiveWork)
					Expect(err).To(BeNil())
					defer resp.Body.Close()
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
				})

				It("Возвращает 404", func() {
					resp, err := http.Get("http://localhost:8101/preventive_works/1")
					Expect(err).To(BeNil())
					defer resp.Body.Close()
					Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
				})
			})

			Context("Проверка возвращаемого json", func() {
				var rightJSON []byte
				BeforeEach(func() {
					_ = json.Unmarshal([]byte("{\"id\": \"\","+
						"\"create_at\": \"2022-01-02T15:04:05.000Z\","+
						"\"deadline\": \"2022-01-04T15:04:05.000Z\","+
						"\"title\": \"test\","+
						"\"description\": \"test\","+
						"\"id_service\": \"62f4eac8f16f2780ae3e637c\","+
						"\"events\": [{"+
						"\"create_at\": \"2022-01-02T15:04:05.000Z\","+
						"\"deadline\": \"2022-01-04T15:04:05.000Z\","+
						"\"description\": \"test\","+
						"\"status\": \"Запланированно\""+
						"}]}"), &rightJSON)
				})
				It("Возвращает json", func() {

					resp, err := http.Get("http://localhost:8101/preventive_works/" + idPreventiveWork)
					Expect(err).To(BeNil())
					defer resp.Body.Close()

					var returnBody []byte
					err = json.NewDecoder(resp.Body).Decode(&returnBody)
					if err != nil {
						return
					}
					Expect(returnBody).To(MatchJSON(rightJSON))
				})

				It("Возвращает 0, когда значение отсутствует", func() {

					resp, err := http.Get("http://localhost:8101/preventive_works/1")
					Expect(err).To(BeNil())
					defer resp.Body.Close()

					var returnBody []byte
					_, err = resp.Body.Read(returnBody)
					if err != nil {
						return
					}
					Expect(returnBody).To(BeZero())
				})
			})
		})
	})

})
