package cmd

import (
	"github.com/adam-hanna/go-oauth2-server/config"
	"github.com/adam-hanna/go-oauth2-server/database"
	"github.com/adam-hanna/go-oauth2-server/health"
	"github.com/adam-hanna/go-oauth2-server/oauth"
	"github.com/adam-hanna/go-oauth2-server/session"
	"github.com/adam-hanna/go-oauth2-server/web"
	"github.com/jinzhu/gorm"
)

var (
	healthService  health.ServiceInterface
	oauthService   oauth.ServiceInterface
	webService     web.ServiceInterface
	sessionService session.ServiceInterface
)

// initConfigDB loads the configuration and connects to the database
func initConfigDB(mustLoadOnce, keepReloading bool, configBackend string) (*config.Config, *gorm.DB, error) {
	// Config
	cnf := config.NewConfig(mustLoadOnce, keepReloading, configBackend)

	// Database
	db, err := database.NewDatabase(cnf)
	if err != nil {
		return nil, nil, err
	}

	return cnf, db, nil
}
