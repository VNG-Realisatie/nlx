module go.nlx.io/nlx

go 1.15

// Override the golang.org/x/text dependency, in version v0.3.3 a vulnerability is fixed.
// See: https://nvd.nist.gov/vuln/detail/CVE-2020-14040.
replace golang.org/x/text => golang.org/x/text v0.3.6

// Override the gopkg.in/yaml.v2 dependency. Versions before v2.2.3 are vulnerable to a Billion Laughs Attack.
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0

require (
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/cloudflare/cfssl v1.5.0
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fgrosse/zaptest v1.1.0
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-errors/errors v1.4.0
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/gorilla/sessions v1.2.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/huandu/xstrings v1.3.2
	github.com/jessevdk/go-flags v1.5.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/jpillora/backoff v1.0.0
	github.com/lib/pq v1.10.2
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210520170846-37e1c6afe023
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/text v0.3.6
	google.golang.org/genproto v0.0.0-20210521181308-5ccab8a35a9a
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/ini.v1 v1.61.0 // indirect
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
	xojoc.pw/useragent v0.0.0-20200116211053-1ec61d55e8fe
)
