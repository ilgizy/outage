using IncidentHistoryService.Models;
using IncidentHistoryService.Models.Interfaces;
using Microsoft.AspNetCore.Mvc;

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
        /// Конструктор по умолчанию<br/>
        /// Создает начальные сущности
        /// </summary>
        public IncidentController()
        {
            if (_storage.Incidents.Count() == 0)
            {
                Incident incident_1 = new(1, "Проблемы с Cloud", "Cloud", new() { "Russia", "German" }, new() { "Serious" });
                incident_1.AddMark(1, "Проблема была обнаружена", DateTimeOffset.Parse("10.07.2022"), "Investigation");
                incident_1.AddMark(2, "Проблема была решена", DateTimeOffset.Parse("12.07.2022"), "Resolved");

                Incident incident_2 = new(2, "Проблемы с DNS", "DNS", new() { "France", "Spain" }, new() { "Small" });
                incident_2.AddMark(3, "Проблема была обнаружена", DateTimeOffset.Parse("11.07.2022"), "Investigation");

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
    }
}
