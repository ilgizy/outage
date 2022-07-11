using IncidentHistoryService.Models;
using IncidentHistoryService.Models.Interfaces;
using Microsoft.AspNetCore.Mvc;
using System.Globalization;

namespace IncidentHistoryService.Controllers
{
    /// <summary>
    /// Контроллер Incident
    /// </summary>
    [Route("api/[controller]")]
    [ApiController]
    public class IncidentController : ControllerBase
    {
        /// <summary>
        /// Внутреннее хранилище
        /// </summary>
        private static IContainable _storage = new VirtualDataSource();

        /// <summary>
        /// Настройка для российского формата даты
        /// </summary>
        private static CultureInfo _cultureInfo = new CultureInfo("ru-RU", false);

        /// <summary>
        /// Конструктор по умолчанию<br/>
        /// Создает начальные сущности
        /// </summary>
        public IncidentController()
        {
            if (_storage.Incidents.Count() == 0)
            {
                Incident incident_1 = new(1, "Проблемы с Cloud", "Cloud", new() { "Russia", "German" }, new() { "Serious" });
                incident_1.AddMark(1, "Проблема была обнаружена", DateTimeOffset.Parse("10.07.2022", _cultureInfo), "Investigation");
                incident_1.AddMark(2, "Проблема была решена", DateTimeOffset.Parse("12.07.2022", _cultureInfo), "Resolved");

                Incident incident_2 = new(2, "Проблемы с DNS", "DNS", new() { "France", "Spain" }, new() { "Small" });
                incident_2.AddMark(3, "Проблема была обнаружена", DateTimeOffset.Parse("11.07.2022", _cultureInfo), "Investigation");

                _storage.Add(incident_1);
                _storage.Add(incident_2);
            }
        }

        /// <summary>
        /// Метод получения списка Incident
        /// </summary>
        /// <returns>JSON со всеми Incident</returns>
        [HttpGet]
        public ActionResult<IEnumerable<Incident>> Get()
        {
            return new ObjectResult(_storage.Incidents);
        }

        /// <summary>
        /// Метод получения конкретного Incident
        /// </summary>
        /// <param name="id">Идентификатор искмого Incident</param>
        /// <returns>Incident с указанным идентификатором</returns>
        [HttpGet("{id}")]
        public ActionResult<IEnumerable<Incident>> Get(int id)
        {
            Incident? incident = _storage.Incidents.FirstOrDefault(x => x.Id == id);
            if (incident == null)
                return NotFound();
            return new ObjectResult(incident);
        }
        
        /// <summary>
        /// Метод создания нового инцидента
        /// </summary>
        /// <param name="id">Идентификатор</param>
        /// <param name="name">Заголовок</param>
        /// <param name="unavailableService">Недоступный во время инцидента сервис</param>
        /// <param name="unavailableZones">Недоступные во время инцидента зоны</param>
        /// <param name="tags">Теги</param>
        /// <returns>
        /// 200 - успешное добавление, возвращает созданный инцидент<br/>
        /// 400 - добавление не удалось
        /// </returns>
        [HttpPost]
        public ActionResult<Incident> Post(int id, string name, string unavailableService,
            [FromQuery] List<string> unavailableZones, [FromQuery] List<string> tags)
        {
            Incident incident = new(id, name, unavailableService, unavailableZones, tags);

            if (_storage.Add(incident))
                return Ok(incident);
            return BadRequest();
        }

        /// <summary>
        /// Метод добавления отметки в историю инцидента
        /// </summary>
        /// <param name="incidentId">Идентификатор инцидента</param>
        /// <param name="markId">Идентификатор метки</param>
        /// <param name="comment">Комментарий метки</param>
        /// <param name="date">Время метки</param>
        /// <param name="tag">Тег метки</param>
        /// <returns>
        /// 200 - успешное добавление, возвращает созданную метку<br/>
        /// 400 - добавление не удалось
        /// </returns>
        [HttpPut]
        public ActionResult<HistoryMark> Put(int incidentId, int markId, string comment, string date, string tag)
        {
            HistoryMark historyMark = new(markId, comment, DateTimeOffset.Parse(date, _cultureInfo), tag, incidentId);
            if (_storage.Add(historyMark))
                return Ok(historyMark);
            return BadRequest();
        }
    }
}
