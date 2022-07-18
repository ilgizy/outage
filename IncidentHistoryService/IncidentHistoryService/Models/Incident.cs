namespace IncidentHistoryService.Models
{
    /// <summary>
    /// Класс инцидента
    /// </summary>
    public class Incident
    {
        /// <summary>
        /// Идентификатор инцидента
        /// </summary>
        public int Id { get; set; }

        /// <summary>
        /// Заголовок инцидента
        /// </summary>
        public string Name { get; set; }

        /// <summary>
        /// Время начала инцидента
        /// </summary>
        /// <returns>
        /// Время первой отметки из HistoryMarks<br/>
        /// Если History не содержит элементов, вернет null
        /// </returns>
        public DateTimeOffset? StartDate
        {
            get
            {
                if (HistoryMarks.Count == 0) return null;
                return HistoryMarks.MinBy(x => x.Date).Date;
            }
        }

        /// <summary>
        /// Время окончания инцидента
        /// </summary>
        /// <returns>
        /// Время последней отметки из HistoryMarks<br/>
        /// Если History не содержит элементов, вернет null
        /// </returns>
        public DateTimeOffset? EndDate
        {
            get
            {
                if (HistoryMarks.Count == 0) return null;
                return HistoryMarks.MaxBy(x => x.Date).Date;
            }
        }

        /// <summary>
        /// Последний комментарий
        /// </summary>
        /// <returns>
        /// Комментарий последней отметки из HistoryMarks<br/>
        /// Если History не содержит элементов, вернет null
        /// </returns>
        public string? LastComment
        {
            get
            {
                if (HistoryMarks.Count == 0) return null;
                return HistoryMarks.MaxBy(x => x.Date).Comment;
            }
        }

        /// <summary>
        /// Недоступный во время инцидента сервис
        /// </summary>
        public string UnavailableService { get; set; }

        /// <summary>
        /// Недоступные во время инцидента зоны
        /// </summary>
        public List<string> UnavailableZones { get; set; }

        /// <summary>
        /// Список тегов
        /// </summary>
        public List<string> Tags { get; set; }

        /// <summary>
        /// Список отметок для ведения истории инцидента
        /// </summary>
        public List<HistoryMark> HistoryMarks { get; set; }

        /// <summary>
        /// Конструктор с параметрами<br/>
        /// Ориентирован на работу с БД
        /// </summary>
        /// <param name="name">Заголовок</param>
        /// <param name="unavailableService">Недоступный во время инцидента сервис</param>
        /// <param name="unavailableZones">Недоступные во время инцидента зоны</param>
        /// <param name="tags">Список тегов</param>
        public Incident(string name, string unavailableService, List<string> unavailableZones, List<string> tags)
        {
            Id = default;
            Name = name;
            UnavailableService = unavailableService;
            UnavailableZones = unavailableZones;
            Tags = tags;
            HistoryMarks = new List<HistoryMark>();
        }
    }
}