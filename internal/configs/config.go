package configs

import (
	"backend-speaker-clone/internal/constants"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	ENV         string = "ENV"
	MODULE_NAME string = "MODULE_NAME"
	MODULE_PORT string = "MODULE_PORT"

	POSTGRES_DB_HOST       string = "POSTGRES_DB_HOST"
	POSTGRES_DB_PORT       string = "POSTGRES_DB_PORT"
	POSTGRES_DB_NAME       string = "POSTGRES_DB_NAME"
	POSTGRES_DB_USER       string = "POSTGRES_DB_USER"
	POSTGRES_DB_PASS       string = "POSTGRES_DB_PASS"
	POSTGRES_DB_SCHEMA     string = "POSTGRES_DB_SCHEMA"
	POSTGRES_DB_IS_DEBUG   string = "POSTGRES_DB_IS_DEBUG"
	POSTGRES_DB_IS_MIGRATE string = "POSTGRES_DB_IS_MIGRATE"

	REDIS_HOST                 string = "REDIS_HOST"
	REDIS_PORT                 string = "REDIS_PORT"
	REDIS_DB                   string = "REDIS_DB"
	REDIS_PASS                 string = "REDIS_PASS"
	REDIS_INSECURE_SKIP_VERIFY string = "REDIS_INSECURE_SKIP_VERIFY"
	REDIS_CUSTOMER_INFO_TTL    string = "REDIS_CUSTOMER_INFO_TTL"
	REDIS_VOICE_CONTEXT_TTL    string = "REDIS_VOICE_CONTEXT_TTL"

	AUTH_REDIS_DB                   string = "AUTH_REDIS_DB"
	AUTH_REDIS_HOST                 string = "AUTH_REDIS_HOST"
	KEYCLOAK_CLIENT_ID              string = "KEYCLOAK_CLIENT_ID"
	KEYCLOAK_CLIENT_SECRET          string = "KEYCLOAK_CLIENT_SECRET"
	KEYCLOAK_HOST                   string = "KEYCLOAK_HOST"
	KEYCLOAK_REALM                  string = "KEYCLOAK_REALM"
	KEYCLOAK_TOKEN_BEARER           string = "KEYCLOAK_TOKEN_BEARER"
	KEYCLOAK_CUSTOMER_CLIENT_ID     string = "KEYCLOAK_CUSTOMER_CLIENT_ID"
	KEYCLOAK_CUSTOMER_CLIENT_SECRET string = "KEYCLOAK_CUSTOMER_CLIENT_SECRET"

	AWS_S3_BUCKET       string = "AWS_S3_BUCKET"
	AWS_S3_REGION       string = "AWS_S3_REGION"
	AWS_S3_STORAGE_PATH string = "AWS_S3_STORAGE_PATH"

	KAFKA_URL      string = "KAFKA_URL"
	KAFKA_TOPIC    string = "KAFKA_TOPIC"
	KAFKA_GROUP_ID string = "KAFKA_GROUP_ID"

	CACHE_BACKEND_SPEAKER_PREFIX string = "CACHE_BACKEND_SPEAKER_PREFIX"

	X_AUTH_API_KEY  string = "X_AUTH_API_KEY"
	ZSS_SERVICE_URL string = "ZSS_SERVICE_URL"

	PROMOTIONS_SERVICE_URL  string = "PROMOTIONS_SERVICE_URL"
	USER_SERVICE_URL        string = "USER_SERVICE_URL"
	CART_SERVICE_URL        string = "CART_SERVICE_URL"
	CATALOG_PRO_SERVICE_URL string = "CATALOG_PRO_SERVICE_URL"
	ASSISTANT_SERVICE_URL   string = "ASSISTANT_SERVICE_URL"

	VOICE_VIRTUAL_ASSISTANT_SERVICE_URL     string = "VOICE_VIRTUAL_ASSISTANT_SERVICE_URL"
	VOICE_VIRTUAL_ASSISTANT_SERVICE_API_KEY string = "VOICE_VIRTUAL_ASSISTANT_SERVICE_API_KEY"

	FIREBASE_API_URL        string = "FIREBASE_API_URL"
	FIREBASE_MEASUREMENT_ID string = "FIREBASE_MEASUREMENT_ID"
)

type environment int

const (
	dev environment = iota
	stg
	prod
)

func LoadEnvFile() {
	envFile := ".env"
	if GetEnv() != "" {
		envFile += "." + GetEnv()
	}
	godotenv.Load(envFile)
}

func GetEnv() string {
	return os.Getenv(ENV)
}

func (e environment) String() string {
	switch e {
	case dev:
		return "dev"
	case stg:
		return "stg"
	case prod:
		return "prod"
	default:
		return "unknown"
	}
}

func getEnvironment() environment {
	env := dev
	e := GetEnv()
	if e == "stg" {
		env = stg
	}
	if e == "prod" {
		env = prod
	}

	return env
}

func IsProduction() bool {
	return getEnvironment() == prod
}

func IsDevelopment() bool {
	return getEnvironment() == dev
}

func GetModuleName() string {
	return os.Getenv(MODULE_NAME)
}
func GetModulePort() string {
	return os.Getenv(MODULE_PORT)
}

// =============== Get postgres config ========================
func GetPostgresHost() string {
	return os.Getenv(POSTGRES_DB_HOST)
}
func GetPostgresPort() string {
	return os.Getenv(POSTGRES_DB_PORT)
}
func GetPostgresName() string {
	return os.Getenv(POSTGRES_DB_NAME)
}
func GetPostgresUser() string {
	return os.Getenv(POSTGRES_DB_USER)
}
func GetPostgresPassword() string {
	return os.Getenv(POSTGRES_DB_PASS)
}
func GetPostgresSchema() string {
	return os.Getenv(POSTGRES_DB_SCHEMA)
}
func GetPostgresDebug() bool {
	return strings.Replace(strings.ToLower(os.Getenv(POSTGRES_DB_IS_DEBUG)), " ", "", -1) == "true"
}
func GetPostgresIsMigrated() bool {
	return strings.Replace(strings.ToLower(os.Getenv(POSTGRES_DB_IS_MIGRATE)), " ", "", -1) == "true"
}

// =============== Get redis config ========================
func GetRedisHost() string {
	return os.Getenv(REDIS_HOST)
}

func GetRedisPort() string {
	return os.Getenv(REDIS_PORT)
}

func GetRedisDb() int {
	redisDb, err := strconv.Atoi(os.Getenv(REDIS_DB))
	if err != nil {
		log.Println(err)
	}
	return redisDb
}

func GetRedisPass() string {
	return os.Getenv(REDIS_PASS)
}

func GetRedisInsecuSkipVerify() bool {
	return strings.Replace(strings.ToLower(os.Getenv(REDIS_INSECURE_SKIP_VERIFY)), " ", "", -1) == "true"
}

func GetRedisCustomerInfoTtl() int64 {
	ttl, err := strconv.ParseInt(os.Getenv(REDIS_CUSTOMER_INFO_TTL), 10, 64)
	if err != nil {
		return constants.DEFAULT_REDIS_CUSTOMER_INFO_TTL
	}
	return ttl
}

func GetRedisVoiceContextTtl() int64 {
	ttl, err := strconv.ParseInt(os.Getenv(REDIS_VOICE_CONTEXT_TTL), 10, 64)
	if err != nil {
		return constants.DEFAULT_REDIS_VOICE_CONTEXT_TTL
	}
	return ttl
}

// =============== Get keycloak config ========================
func GetAuthRedisDb() int {
	redisDb, _ := strconv.Atoi(os.Getenv(AUTH_REDIS_DB))
	return redisDb
}
func GetAuthRedisHost() string {
	return os.Getenv(AUTH_REDIS_HOST)
}
func GetKeycloakHost() string {
	return os.Getenv(KEYCLOAK_HOST)
}
func GetKeycloakRealm() string {
	return os.Getenv(KEYCLOAK_REALM)
}
func GetKeycloakClientId() string {
	return os.Getenv(KEYCLOAK_CLIENT_ID)
}
func GetKeycloakClientSecret() string {
	return os.Getenv(KEYCLOAK_CLIENT_SECRET)
}
func GetKeycloakTokenBearer() string {
	return os.Getenv(KEYCLOAK_TOKEN_BEARER)
}

func GetKeycloakCustomerClientId() string {
	return os.Getenv(KEYCLOAK_CUSTOMER_CLIENT_ID)
}

func GetKeycloakCustomerClientSecret() string {
	return os.Getenv(KEYCLOAK_CUSTOMER_CLIENT_SECRET)
}

// =============== Kafka ========================
func GetKafkaUrl() string {
	return os.Getenv(KAFKA_URL)
}
func GetKafkaTopic() string {
	return os.Getenv(KAFKA_TOPIC)
}
func GetKafkaGroupId() string {
	return os.Getenv(KAFKA_GROUP_ID)
}

// =============== S3 ========================
func GetS3Bucket() string {
	return os.Getenv(AWS_S3_BUCKET)
}
func GetS3Region() string {
	return os.Getenv(AWS_S3_REGION)
}
func GetS3StoragePath() string {
	return os.Getenv(AWS_S3_STORAGE_PATH)
}

// ============ Cache ==========================
func GetCacheBackendSpeakerPrefix() string {
	return os.Getenv(CACHE_BACKEND_SPEAKER_PREFIX)
}

// =============== Other =====================
func GetAuthApiKey() string {
	return os.Getenv(X_AUTH_API_KEY)
}

func GetZssEndpoint() string {
	return os.Getenv(ZSS_SERVICE_URL)
}

// =============== Services ========================
func GetPromotionsServiceUrl() string {
	return os.Getenv(PROMOTIONS_SERVICE_URL)
}

func GetUserCustomerEndpoint() string {
	return os.Getenv(USER_SERVICE_URL)
}

func GetCartServiceUrl() string {
	return os.Getenv(CART_SERVICE_URL)
}

func GetCatalogProEndpoint() string {
	return os.Getenv(CATALOG_PRO_SERVICE_URL)
}

func GetAssistantService() string {
	return os.Getenv(ASSISTANT_SERVICE_URL)
}

func GetFirebaseApiUrl() string {
	return os.Getenv(FIREBASE_API_URL)
}

func GetFirebaseMeasureement() string {
	return os.Getenv(FIREBASE_MEASUREMENT_ID)
}
