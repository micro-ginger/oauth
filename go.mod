module github.com/micro-ginger/oauth

go 1.22.3

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/gin-gonic/gin v1.9.1
	github.com/ginger-core/compound v0.0.0-20230608151919-2963b75416c3
	github.com/ginger-core/compound/registry v0.0.0-20240815151007-6306bf13a816
	github.com/ginger-core/errors v0.0.0-20230703084505-b10c3f9cedfb
	github.com/ginger-core/gateway v0.0.0-20240909095814-36fbbeaa9104
	github.com/ginger-core/log v0.0.0-20240629145652-3b2876535940
	github.com/ginger-core/query v0.0.0-20230608153800-9375f70642d8
	github.com/ginger-core/repository v0.0.0-20230608165607-87044af67011
	github.com/ginger-gateway/ginger v0.0.0-20240909095921-d32610dbcaf4
	github.com/ginger-repository/redis v0.0.0-20230608170101-0b74d866bc2d
	github.com/ginger-repository/sql v0.0.0-20240806160111-7365d88b4d19
	github.com/go-sql-driver/mysql v1.7.1
	github.com/micro-blonde/auth v0.0.0-20241004080911-a8f21e8940d6
	github.com/micro-blonde/auth/authorization v0.0.0-20230630082657-5b527d26afce
	github.com/micro-blonde/auth/proto v0.0.0-20241004080911-a8f21e8940d6
	github.com/micro-blonde/file v0.0.0-20240805212943-31734ed43a25
	github.com/micro-blonde/file/client v0.0.0-20240805212943-31734ed43a25
	github.com/nicksnyder/go-i18n/v2 v2.2.1
	golang.org/x/crypto v0.23.0
	golang.org/x/text v0.15.0
	google.golang.org/protobuf v1.34.2
	gorm.io/gorm v1.25.1
)

require (
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/ginger-core/errors/grpc v0.0.0-20230703084505-b10c3f9cedfb // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/micro-blonde/file/proto v0.0.0-20240805211729-18bd31844877 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.16.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240624140628-dc46fd24d27d // indirect
	google.golang.org/grpc v1.65.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/driver/postgres v1.5.2 // indirect
)

// replace github.com/micro-blonde/auth => ../../micro-blonde/auth

// replace github.com/ginger-repository/sql => ../../../github.com/ginger-repository/sql
// replace github.com/micro-blonde/auth/proto => ../../micro-blonde/auth/proto
