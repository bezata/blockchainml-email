package services

type Services struct {
    Email *EmailService
    Staff *StaffService
    Auth  *AuthService
}

func New(cfg Config) *Services {
    return &Services{
        Email: NewEmailService(EmailServiceConfig{
            Repo:     cfg.Repositories.Email,
            Cache:    cfg.Cache,
            Search:   cfg.Search,
            Notifier: cfg.Notifier,
            Logger:   cfg.Logger,
            Metrics:  cfg.Metrics,
        }),
        Staff: NewStaffService(StaffServiceConfig{
            Repo:    cfg.Repositories.Staff,
            Cache:   cfg.Cache,
            Logger:  cfg.Logger,
            Metrics: cfg.Metrics,
        }),
        Auth: NewAuthService(AuthServiceConfig{
            Repo:    cfg.Repositories.Staff,
            Config:  cfg.Config.JWT,
            Logger:  cfg.Logger,
            Metrics: cfg.Metrics,
        }),
    }
}