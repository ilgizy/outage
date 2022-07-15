using Microsoft.EntityFrameworkCore;

namespace IncidentHistoryService.Models
{
    public class ApplicationContext : DbContext
    {
        public DbSet<Incident> Incidents { get; set; } = null!;
        public DbSet<HistoryMark> HistoryMarks { get; set; } = null!;

        public ApplicationContext()
        {
            AppContext.SetSwitch("Npgsql.EnableLegacyTimestampBehavior", true);
            Database.EnsureCreated();
        }

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            optionsBuilder.EnableSensitiveDataLogging();
            optionsBuilder.UseNpgsql("Host=localhost;Port=5432;Database=IncidentHistoryServiceDB;Username=postgres;Password=2458173671;Include Error Detail=true");
        }
    }
}
