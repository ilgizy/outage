using IncidentHistoryService.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
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
        /// Настройка для российского формата даты
        /// </summary>
        private static CultureInfo _cultureInfo = new CultureInfo("ru-RU", false);

        /// <summary>
        /// Конструктор по умолчанию<br/>
        /// Создает начальные сущности
        /// </summary>
        public IncidentController()
        {
            using ApplicationContext db = new();
            if (!db.Incidents.Any())
            {
                Incident incident_1 = new("Проблемы с Cloud", "Cloud", new() { "Russia", "German" }, new() { "Serious" });
                HistoryMark mark_1 = new("Проблема была обнаружена", DateTimeOffset.Parse("10.07.2022", _cultureInfo), "Investigation", incident_1);
                HistoryMark mark_2 = new("Проблема была решена", DateTimeOffset.Parse("12.07.2022", _cultureInfo), "Resolved", incident_1);

                Incident incident_2 = new("Проблемы с DNS", "DNS", new() { "France", "Spain" }, new() { "Small" });
                HistoryMark mark_3 = new("Проблема была обнаружена", DateTimeOffset.Parse("11.07.2022", _cultureInfo), "Investigation", incident_2);

                db.Incidents.Add(incident_1);
                db.Incidents.Add(incident_2);
                db.HistoryMarks.Add(mark_1);
                db.HistoryMarks.Add(mark_2);
                db.HistoryMarks.Add(mark_3);
                db.SaveChanges();
            }
        }

        /// <summary>
        /// Метод получения списка Incident
        /// </summary>
        /// <returns>JSON со всеми Incident</returns>
        [HttpGet]
        public ActionResult<IEnumerable<Incident>> Get()
        {
            using ApplicationContext db = new();
            return new ObjectResult(db.Incidents.Include(x => x.HistoryMarks).ToList());
        }

        /// <summary>
        /// Метод получения конкретного Incident
        /// </summary>
        /// <param name="id">Идентификатор искмого Incident</param>
        /// <returns>Incident с указанным идентификатором</returns>
        [HttpGet("{id}")]
        public ActionResult<IEnumerable<Incident>> Get(int id)
        {
            using ApplicationContext db = new();
            Incident? incident = db.Incidents.Include(x => x.HistoryMarks).FirstOrDefault(x => x.Id == id);
            if (incident == null)
                return NotFound();
            return new ObjectResult(incident);
        }
        
        /// <summary>
        /// Метод создания нового инцидента
        /// </summary>
        /// <param name="name">Заголовок</param>
        /// <param name="unavailableService">Недоступный во время инцидента сервис</param>
        /// <param name="unavailableZones">Недоступные во время инцидента зоны</param>
        /// <param name="tags">Теги</param>
        /// <returns>
        /// 200 - успешное добавление, возвращает созданный инцидент<br/>
        /// </returns>
        [HttpPost]
        public ActionResult<Incident> Post(string name, string unavailableService,
            [FromQuery] List<string> unavailableZones, [FromQuery] List<string> tags)
        {
            Incident incident = new(name, unavailableService, unavailableZones, tags);

            using ApplicationContext db = new();
            db.Incidents.Add(incident);
            db.SaveChanges();

            return Ok(incident);
        }

        /// <summary>
        /// Метод добавления отметки в историю инцидента
        /// </summary>
        /// <param name="incidentId">Идентификатор инцидента</param>
        /// <param name="comment">Комментарий метки</param>
        /// <param name="date">Время метки</param>
        /// <param name="tag">Тег метки</param>
        /// <returns>
        /// 200 - успешное добавление, возвращает созданную метку<br/>
        /// 400 - добавление не удалось
        /// </returns>
        [HttpPut]
        public ActionResult<HistoryMark> Put(int incidentId, string comment, string date, string tag)
        {
            using ApplicationContext db = new();

            Incident? incident = db.Incidents.Include(x => x.HistoryMarks).FirstOrDefault(x => x.Id == incidentId);
            if (incident == null)
                return BadRequest();

            HistoryMark historyMark = new(comment, DateTimeOffset.Parse(date, _cultureInfo), tag, incident);

            incident.HistoryMarks.Add(historyMark);
            db.Update(incident);
            db.SaveChanges();

            return Ok(historyMark);
        }
    }
}
