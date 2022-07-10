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
        /// Время первой отметки из History<br/>
        /// Если History не содержит элементов, вернет null
        /// </returns>
        public DateTimeOffset? StartDate
        {
            get
            {
                if (History.Count == 0) return null;
                return History.MinBy(x => x.Date).Date;
            }
        }

        /// <summary>
        /// Время окончания инцидента
        /// </summary>
        /// <returns>
        /// Время последней отметки из History<br/>
        /// Если History не содержит элементов, вернет null
        /// </returns>
        public DateTimeOffset? EndDate
        {
            get
            {
                if (History.Count == 0) return null;
                return History.MaxBy(x => x.Date).Date;
            }
        }

        /// <summary>
        /// Последний комментарий
        /// </summary>
        /// <returns>
        /// Комментарий последней отметки из History<br/>
        /// Если History не содержит элементов, вернет null
        /// </returns>
        public string? LastComment
        {
            get
            {
                if (History.Count == 0) return null;
                return History.MaxBy(x => x.Date).Comment;
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
        public List<HistoryMark> History { get; set; }

        /// <summary>
        /// Конструктор по умолчанию<br/>
        /// Инициализирует списки<br/>
        /// - UnavailableZones<br/>
        /// - Tags<br/>
        /// - History
        /// </summary>
        public Incident()
        {
            UnavailableZones = new List<string>();
            Tags = new List<string>();
            History = new List<HistoryMark>();
        }

        /// <summary>
        /// Добавление отметки в истории
        /// </summary>
        /// <param name="id">Идентификатор отметки</param>
        /// <param name="comment">Комментарий отметки</param>
        /// <param name="date">Время отметки</param>
        /// <param name="tag">Тег отметки</param>
        public void AddMark(int id, string comment, DateTimeOffset date, string tag)
        {
            History.Add(new HistoryMark(id, comment, date, tag, this));
        }
    }
}