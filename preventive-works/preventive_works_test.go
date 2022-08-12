package preventive_works_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/url"
	"strings"
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

				bodyBytes, err := io.ReadAll(resp.Body)
				idPreventiveWork = string(bodyBytes)
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

			It("Возвращает 400, если дата введена неверно", func() {
				resp, err := http.PostForm("http://localhost:8101/preventive_works/new_work", url.Values{
					"name_service": {"test"},
					"create_at":    {"рп"},
					"deadline":     {"2022-01-02 15:04:05"},
					"title":        {"test"},
					"description":  {"test"}})
				Expect(err).To(BeNil())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
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
				var rightJSON string
				BeforeEach(func() {
					rightJSON = "{\"id\": \"\"," +
						"\"create_at\": \"2022-01-02T15:04:05Z\"," +
						"\"deadline\": \"2022-01-04T15:04:05Z\"," +
						"\"title\": \"test\"," +
						"\"description\": \"test\"," +
						"\"id_service\": \"62f4eac8f16f2780ae3e637c\"," +
						"\"events\": [{" +
						"\"create_at\": \"2022-01-02T15:04:05Z\"," +
						"\"deadline\": \"2022-01-04T15:04:05Z\"," +
						"\"description\": \"test\"," +
						"\"status\": \"Запланированно\"" +
						"}]}"
				})
				It("Возвращает json", func() {

					resp, err := http.Get("http://localhost:8101/preventive_works/" + idPreventiveWork)
					Expect(err).To(BeNil())
					defer resp.Body.Close()

					fmt.Println(idPreventiveWork)
					returnBody, err := io.ReadAll(resp.Body)
					Expect(err).To(BeNil())

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

	Describe("Проверка preventive_works/{id}/new_event", func() {
		Context("Проверка возвращаемых кодов", func() {
			var client *http.Client
			BeforeEach(func() {
				client = &http.Client{}
			})
			It("Возвращает 200", func() {
				data := url.Values{
					"status":      {"update"},
					"create_at":   {"2006-01-02 15:04:05"},
					"deadline":    {"2006-01-03 15:04:05"},
					"description": {"test"}}

				r, err := http.NewRequest(http.MethodPut, "http://localhost:8101/preventive_works/62f6755176708770fec2015f/new_event", strings.NewReader(data.Encode()))
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				Expect(err).To(BeNil())

				resp, err := client.Do(r)
				Expect(err).To(BeNil())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})

			It("Возвращает 500, если дата окончания раньше даты создания", func() {
				data := url.Values{
					"status":      {"update"},
					"create_at":   {"2006-01-02 15:04:05"},
					"deadline":    {"2006-01-01 15:04:05"},
					"description": {"test"}}

				r, err := http.NewRequest(http.MethodPut, "http://localhost:8101/preventive_works/62f6755176708770fec2015f/new_event", strings.NewReader(data.Encode()))
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				Expect(err).To(BeNil())

				resp, err := client.Do(r)
				Expect(err).To(BeNil())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			})

			It("Возвращает 500, если дата введена неверно", func() {
				data := url.Values{
					"status":      {"update"},
					"create_at":   {"2006-01-02 15:04:05"},
					"deadline":    {"лорп"},
					"description": {"test"}}

				r, err := http.NewRequest(http.MethodPut, "http://localhost:8101/preventive_works/62f6755176708770fec2015f/new_event", strings.NewReader(data.Encode()))
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				Expect(err).To(BeNil())

				resp, err := client.Do(r)
				Expect(err).To(BeNil())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
