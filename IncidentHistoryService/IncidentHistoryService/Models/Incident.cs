namespace IncidentHistoryService.Models
{
    public class Incident
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public DateTimeOffset? StartDate
        {
            get
            {
                if (History.Count == 0) return null;
                return History[0].Date;
            }
        }
        public DateTimeOffset? EndDate
        {
            get
            {
                if (History.Count == 0) return null;
                return History.Last().Date;
            }
        }
        public string? LastComment
        {
            get
            {
                if (History.Count == 0) return null;
                return History[0].Comment;
            }
        }
        public string UnavailableService { get; set; }
        public string UnavailableZone { get; set; }
        public List<string> Tags { get; set; }

        public List<HistoryMark> History { get; set; }
        public Incident()
        {
            Tags = new List<string>();
            History = new List<HistoryMark>();
        }
    }
}
