namespace IncidentHistoryService.Models
{
    public class HistoryMark
    {
        public int Id { get; set; }
        public string Comment { get; set; }
        public DateTimeOffset Date { get; set; }
        public string Tag { get; set; }

        public int IncidentId { get; set; }
        public Incident Incident { get; set; }
    }
}
