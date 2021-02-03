module go.nlx.io/nlx

go 1.15

// Override the golang.org/x/text dependency, in version v0.3.3 a vulnerability is fixed.
// See: https://nvd.nist.gov/vuln/detail/CVE-2020-14040.
replace golang.org/x/text => golang.org/x/text v0.3.4

// Override the gopkg.in/yaml.v2 dependency. Versions before v2.2.3 are vulnerable to a Billion Laughs Attack.
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0

// Override the google.golang.org/grpc dependency. One of the package has `@latest` which doesn't work.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/bwesterb/go-atum v1.0.3 // indirect
	github.com/cloudflare/cfssl v1.5.0
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fgrosse/zaptest v1.1.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-errors/errors v1.1.1
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.1
	github.com/gorilla/schema v1.2.0
	github.com/gorilla/sessions v1.2.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.8
	github.com/hashicorp/go-retryablehttp v0.6.4 // indirect
	github.com/huandu/xstrings v1.3.2
	github.com/jasonlvhit/gocron v0.0.0-20200319232826-9b00bf4b9ebc // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/jpillora/backoff v1.0.0
	github.com/ktr0731/toml v0.3.0
	github.com/lib/pq v1.8.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/privacybydesign/irmago v0.6.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.7
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	google.golang.org/genproto v0.0.0-20210202153253-cf70463f6119
	google.golang.org/grpc v1.33.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.8
)
