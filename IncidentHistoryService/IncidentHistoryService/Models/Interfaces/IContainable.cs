namespace IncidentHistoryService.Models.Interfaces
{
    /// <summary>
    /// Интерфейс хранилища Incident и HistoryMark
    /// </summary>
    public interface IContainable
    {
        /// <summary>
        /// Хранилище Incident
        /// </summary>
        public IEnumerable<Incident> Incidents { get; }

        /// <summary>
        /// Хранилище HistoryMark
        /// </summary>
        public IEnumerable<HistoryMark> HistoryMarks { get; }

        /// <summary>
        /// Добавление в хранилище Incident
        /// </summary>
        /// <param name="incident">Новый Incident</param>
        /// <returns>Произошло ли добавление</returns>
        public bool Add(Incident incident);

        /// <summary>
        /// Добавление в хранилище HistoryMark
        /// </summary>
        /// <param name="historyMark">Новый HistoryMark</param>
        /// <returns>Произошло ли добавление</returns>
        public bool Add(HistoryMark historyMark);
    }
}