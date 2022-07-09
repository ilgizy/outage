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
                return History[0].Date;
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
                return History.Last().Date;
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
                return History[0].Comment;
            }
        }

        /// <summary>
        /// Недоступный во время инцидента сервис
        /// </summary>
        public string UnavailableService { get; set; }

        /// <summary>
        /// Недоступная во время инцидента зона
        /// </summary>
        public string UnavailableZone { get; set; }

        /// <summary>
        /// Список тегов
        /// </summary>
        public List<string> Tags { get; set; }

        /// <summary>
        /// Список отметок для ведения истории инцидента
        /// </summary>
        public List<HistoryMark> History { get; set; }

        public Incident()
        {
            Tags = new List<string>();
            History = new List<HistoryMark>();
        }
    }
}
