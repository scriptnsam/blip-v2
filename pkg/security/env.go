package security

import "os"

func LoadEnv() {
	os.Setenv("MYSQL_HOST", "sql.freedb.tech")
	os.Setenv("MYSQL_USERNAME", "freedb_blip_admin")
	os.Setenv("MYSQL_PASSWORD", "*NtbhBjSKPF8wZn")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("DB_NAME", "freedb_blipDb")

	// JWT SECRET KEY TOKEN
	os.Setenv("JWT_SECRET_KEY", "oYO87Vt87OVT67FVT7fv87TV76diuyf67R6967vo78T78OTUY")
}
