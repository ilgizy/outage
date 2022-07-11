﻿using IncidentHistoryService.Models.Interfaces;

namespace IncidentHistoryService.Models
{
    /// <summary>
    /// Виртуальный источник Incident и HistoryMark
    /// </summary>
    public class VirtualDataSource : IContainable
    {
        /// <summary>
        /// Список Incident
        /// </summary>
        private List<Incident> _incidents;

        /// <summary>
        /// Список HistoryMark
        /// </summary>
        private List<HistoryMark> _historyMarks;

        public IEnumerable<Incident> Incidents
        {
            get => _incidents;
        }

        public IEnumerable<HistoryMark> HistoryMarks
        {
            get => _historyMarks;
        }

        /// <summary>
        /// Конструктор по умолчанию, инициализирующий внутренние списки
        /// </summary>
        public VirtualDataSource()
        {
            _incidents = new List<Incident>();
            _historyMarks = new List<HistoryMark>();
        }

        public bool Add(Incident incident)
        {
            _incidents.Add(incident);
            return true;
        }

        public bool Add(HistoryMark historyMark)
        {
            int index = _incidents.FindIndex(x => x.Id == historyMark.IncidentId);
            if (index >= 0)
            {
                _incidents[index].AddMark(historyMark.Id, historyMark.Comment, historyMark.Date, historyMark.Tag);
                _historyMarks.Add(historyMark);
                return true;
            }
            return false;
        }
    }
}
