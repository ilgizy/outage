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

            var configuration = new ConfigurationBuilder().AddJsonFile("appsettings.json").AddEnvironmentVariables().Build();

            optionsBuilder.UseNpgsql(configuration.GetConnectionString("DefaultConnection"));
        }
    }
}
