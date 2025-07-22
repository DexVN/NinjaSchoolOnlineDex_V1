source .env

if [ -z "$1" ]; then
  echo "❌ Thiếu lệnh Goose! Ví dụ: ./migrate.sh up"
  exit 1
fi

goose -dir migrations postgres "$DB_URL" "$1"