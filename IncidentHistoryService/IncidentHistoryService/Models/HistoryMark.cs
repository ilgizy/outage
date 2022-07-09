namespace IncidentHistoryService.Models
{
    /// <summary>
    /// Отметка в истории инцидента
    /// </summary>
    public class HistoryMark
    {
        /// <summary>
        /// Идентификатор отметки
        /// </summary>
        public int Id { get; set; }

        /// <summary>
        /// Комментарий отметки
        /// </summary>
        public string Comment { get; set; }

        /// <summary>
        /// Время отметки
        /// </summary>
        public DateTimeOffset Date { get; set; }

        /// <summary>
        /// Тег отметки
        /// </summary>
        public string Tag { get; set; }

        /// <summary>
        /// Идентификатор инцидента, к которому относится отметка
        /// </summary>
        public int IncidentId { get; set; }

        /// <summary>
        /// Инцидент, к которому относится отметка
        /// </summary>
        public Incident Incident { get; set; }
    }
}
