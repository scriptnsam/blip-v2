package security

import "os"

func LoadEnv() {
	os.Setenv("MYSQL_HOST", "sql8.freesqldatabase.com")
	os.Setenv("MYSQL_USERNAME", "sql8787889")
	os.Setenv("MYSQL_PASSWORD", "8I6erLRIMf")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("DB_NAME", "sql8787889")

	// JWT SECRET KEY TOKEN
	os.Setenv("JWT_SECRET_KEY", "oYO87Vt87OVT67FVT7fv87TV76diuyf67R6967vo78T78OTUY")
}
